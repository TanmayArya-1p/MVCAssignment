import { markAsPaid, resolveBill } from "../api/orders";
import { useEffect, useState } from "react";
import { orderColourMap } from "../utils/const";
import toast, { Toaster } from "react-hot-toast";
import { useNavigate } from "react-router-dom";



export default function Bill({order, setBillLoading, billLoading, role}) {
    const navigate = useNavigate();
    const [orderItems, setOrderItems] = useState([]);
    const [billableAmount, setBillableAmount] = useState(0);
    const [amountPaid, setAmountPaid] = useState(0);

    const billHandler = async () => {
        if(!window.confirm("Are you sure you want to bill this order? Any pending items will no longer be processed.")) {
            return;
        }
        await resolveBill(order.id, true)
        toast.success("Order billed successfully");
        navigate(0)
    }

    const payHandler = async (orderId) => {
        const paidAmount = document.getElementById("paid-amount").value;
        if(!paidAmount || isNaN(paidAmount) || paidAmount <= 0) {
            alert("Please enter a valid amount.");
            return;
        }
        if(!window.confirm(`Are you sure you want to mark this order as paid with ₹${paidAmount}?`)) {
            return;
        }
        if(paidAmount < billableAmount) {
            toast.error(`Paid amount ₹${paidAmount} is less than the billable amount ₹${billableAmount}. Please enter a valid amount.`);
            return;
        }
        try {
            await markAsPaid(orderId, paidAmount);
            navigate(0);

        } catch (error) {
            toast.error("Failed to mark order as paid.");
        }
    }

    useEffect(() => {
        const fetchBill = async () => {
            setBillLoading(true);
            try {
                const bill = await resolveBill(order.id);
                setOrderItems(bill.items);
                setBillableAmount(bill.billable_amount);
            } catch (error) {
                console.error("error fetching order items:", error);
                navigate("/notfound");
            } finally {
                setBillLoading(false);
            }
        };
        fetchBill();
    }, [order.id, setBillLoading]);

    if (billLoading) {
        return <></>
    }

    return (<div>
        
            <div className="relative overflow-x-auto mt-2 w-fit bg-white rounded-xl shadow-md border">
            <table className="w-fit bg-white p-2">
                <thead className="text-md ubuntu-bold">
                    <tr>
                        <th scope="col" className="px-2 py-3">
                            Item
                        </th>
                        <th scope="col" className="px-2 py-3">
                            Quantity
                        </th>
                        <th scope="col" className="px-2 py-3">
                            Price per Item
                        </th>
                        <th scope="col" className="px-2 py-3">
                            Total Price
                        </th>
                        <th scope="col" className="px-2 py-3">
                            Status
                        </th>
                    </tr>
                </thead>
                <tbody className="text-md text-center ubuntu-regular">
                    {orderItems.map((item,id) => 
                        <tr key={id} className="bg-white border-b">
                            <td className="px-6 py-4">
                                {item.name}
                            </td>
                            <td className="px-6 py-4">
                                {item.quantity}
                            </td>
                            <td className="px-6 py-4">
                                ₹{item.price}
                            </td>
                            <td className="px-6 py-4">
                                ₹{item.price*item.quantity}<br />
                            </td>
                            <td className="px-6 py-4 ubuntu-bold">
                                <span className={`text-${orderColourMap[item.status]}`}>
                                    {item.status}
                                </span>
                            </td>
                        </tr>
                    )}

                    <tr className="bg-gray-100 ubuntu-bold">
                        <td></td>
                        <td></td>                            
                        <td className="px-6 py-4 text-right">Total</td>
                        <td className="px-6 py-4">₹{billableAmount}</td>
                        <td className="px-2">
                            {order.status!=="billed" && order.status!=="paid" && (
                                <button className="" onClick={() => billHandler(order.id)}>Bill Order</button>
                            )}
                            {order.status==="billed" && (role==="chef" || role==="admin") && (
                                <>
                                    <button className="" onClick={() => payHandler(order.id)}>Mark as Paid</button>
                                    <input type="number" id="paid-amount" value={amountPaid}
                                    placeholder="Amount Paid" 
                                    onChange={(e) => setAmountPaid(e.target.value >= 0 ? e.target.value : 0)}
                                    className="mx-5 w-40 focus:outline-0 bg-white p-2 rounded-sm border-2" />
                                </>
                            )}
                            {order.status==="paid" && (
                                <span className="bg-white border-2 rounded-sm p-2">Tip: ₹{order.tip}</span>
                            )}
                        </td>
                    </tr>
                </tbody>
            </table>

        </div>
    </div>)

}
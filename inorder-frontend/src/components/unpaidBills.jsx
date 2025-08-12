import { useEffect, useState } from "react"


export default function UnpaidBills({orders}) {

    const [unpaidOrders, setUnpaidOrders] = useState([]);


    useEffect(()=> {
        setUnpaidOrders(orders.filter(order => order.status === "billed"));
    }, [orders]);



    return <>
        <div id="unpaid-bills-container" className="overflow-x-auto max-w-fit p-2 bg-white rounded-xl shadow-md border-2" style={{ maxHeight: "40rem" }}>
            <table>
                <thead>
                    <tr className="text-center text-lg">
                        <th className="ubuntu-bold px-3 text-left">Order</th>
                        <th className="ubuntu-bold px-3 text-center">Amount</th>
                    </tr>
                </thead>
                <tbody className="text-md ubuntu-regular p-2">
                    {unpaidOrders.map(order => (
                        <tr
                            className="ubuntu-regular text-left p-2 m-10"
                            key={order.id}
                        >
                            <td className="px-3">
                                <a className="order-link ubuntu-bold w-full" onClick={() => window.location.href = `/order/${order.id}`}>
                                    Order #{order.id}
                                </a>
                            </td>
                            <td className="px-3 text-lg font-bold ubuntu-bold text-center">
                                ₹ {order.billable_amount}
                            </td>
                        </tr>
                    ))}
                </tbody>
            </table>
        </div>
    </>
}
    
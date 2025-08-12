import { useEffect, useState } from "react";
import axios from "axios";
import OrderCard from "../../components/orderCard";
import Spinner from "../../components/spinner";
import { getAllOrders, getMyOrders } from "../../api/orders";
import OrderBook from "../../components/orderBook";
import ItemMenu from "../../components/itemMenu";
import CreateOrderModal from "../../components/createOrderModal";
import UnpaidBills from "../../components/unpaidBills";
import { ItemQueue } from "../../components/itemQueue";

export default function AdminHomeScreen() {
    const [ordersLoading, setOrdersLoading] = useState(true);
    const [orders, setOrders] = useState([]);
    const [createModelOpen, setCreateModelOpen] = useState(false);
    const [itemOrders, setItemOrders] = useState({});
    
    useEffect(() => {
        const fetchOrders = async () => {
            try {
                setOrdersLoading(true);
                const response = await getAllOrders(null, 0);
                console.log("orders:", response);
                setOrders(response);
                setOrdersLoading(false);

            }
            catch (error) {
                console.error("error fetching orders:", error);
                setOrdersLoading(false);

            }
        };
        fetchOrders();
    }, []);

    
    return <>
        <div className="flex flex-col mt-7 p-5 mb-20">
            <div className="mb-10 flex flex-row flex-wrap">
                <div className="w-220">
                    <div className="flex flex-row justify-between items-center">
                        <div className="text-3xl ubuntu-bold flex flex-row gap-5 items-center">All Orders
                        <button className="flex flex-row gap-2 text-lg items-center ubuntu-regular text-black" onClick={() => setCreateModelOpen(true)}>
                            <svg xmlns="http://www.w3.org/2000/svg" style={{width:20,height:20}} viewBox="0 0 640 640"><path d="M352 128C352 110.3 337.7 96 320 96C302.3 96 288 110.3 288 128L288 288L128 288C110.3 288 96 302.3 96 320C96 337.7 110.3 352 128 352L288 352L288 512C288 529.7 302.3 544 320 544C337.7 544 352 529.7 352 512L352 352L512 352C529.7 352 544 337.7 544 320C544 302.3 529.7 288 512 288L352 288L352 128z"/></svg>
                            <div className="text-center flex items-center">Create Order</div>
                        </button>
                    </div>
                    {ordersLoading && <Spinner />}
                    </div>
                    <OrderBook orders={orders.filter(order => order.status !== "paid")} setOrders={setOrders} loading={ordersLoading} admin={true}/>

                </div>

                <div className="flex flex-col gap-3">
                    <div id="navigate-to-order" className="ubuntu-bold mt-5 p-2 w-fit flex flex-row gap-2 items-center bg-white shadow-md border-2 rounded">
                        Navigate to Order No:
                        <input type="number" id="navigate-order-id" placeholder="Order No" className="p-1 border-2 w-30 mx-2 rounded-sm"/>
                        <button onClick={() => window.location.href = `/order/${document.getElementById("navigate-order-id").value}`}>
                            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth="1.5" stroke="currentColor" className="size-5">
                                <path strokeLinecap="round" strokeLinejoin="round" d="M6 12 3.269 3.125A59.769 59.769 0 0 1 21.485 12 59.768 59.768 0 0 1 3.27 20.875L5.999 12Zm0 0h7.5" />
                            </svg>
                        </button>
                    </div>

                    <div className="mt-10">
                        <h2 className="text-3xl ubuntu-bold mb-2">Unpaid Bills</h2>
                        <UnpaidBills orders={orders} />

                    </div>


                </div>
            </div>
            <div className="flex flex-col">
                <div className="text-3xl ubuntu-bold mb-2">Previous Orders</div>
                <OrderBook noFilter orders={orders.filter(order => order.status === "paid")} setOrders={setOrders} loading={ordersLoading} admin={true}/>
            </div>




            <CreateOrderModal isOpen={createModelOpen} setIsOpen={setCreateModelOpen} />
            
        </div>

    </>
}

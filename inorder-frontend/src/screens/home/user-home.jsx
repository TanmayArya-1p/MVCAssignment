import { useEffect, useState } from "react";
import axios from "axios";
import OrderCard from "../../components/orderCard";
import Spinner from "../../components/spinner";
import { getMyOrders } from "../../api/orders";
import OrderBook from "../../components/orderBook";
import ItemMenu from "../../components/itemMenu";
import CreateOrderModal from "../../components/createOrderModal";

export default function UserHomeScreen() {
    const [ordersLoading, setOrdersLoading] = useState(true);
    const [orders, setOrders] = useState([]);
    const [createModelOpen, setCreateModelOpen] = useState(false);
    const [itemOrders, setItemOrders] = useState({});
    
    useEffect(() => {
        const fetchOrders = async () => {
            try {
                setOrdersLoading(true);
                const response = await getMyOrders(null, 0);
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
            <div>
                <div className="text-3xl ubuntu-bold flex flex-row gap-5">Your Orders
                <button className="flex flex-row gap-2 text-lg items-center ubuntu-regular text-black" onClick={() => setCreateModelOpen(true)}>
                    <svg xmlns="http://www.w3.org/2000/svg" style={{width:20,height:20}} viewBox="0 0 640 640"><path d="M352 128C352 110.3 337.7 96 320 96C302.3 96 288 110.3 288 128L288 288L128 288C110.3 288 96 302.3 96 320C96 337.7 110.3 352 128 352L288 352L288 512C288 529.7 302.3 544 320 544C337.7 544 352 529.7 352 512L352 352L512 352C529.7 352 544 337.7 544 320C544 302.3 529.7 288 512 288L352 288L352 128z"/></svg>
                    <div>Create Order</div>
                </button>
            </div>
            {ordersLoading && <Spinner />}
            </div>
            <OrderBook orders={orders} loading={ordersLoading}/>


            <div className="text-3xl ubuntu-bold flex flex-row gap-5 mt-10">Item Menu</div>

            <ItemMenu/>
            <CreateOrderModal isOpen={createModelOpen} setIsOpen={setCreateModelOpen} />

        </div>

    </>
}

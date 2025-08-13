import { useEffect, useState } from "react";
import axios from "axios";
import OrderCard from "../../components/orderCard";
import Spinner from "../../components/spinner";
import { getMyOrders } from "../../api/orders";
import OrderBook from "../../components/orderBook";
import ItemMenu from "../../components/itemMenu";
import CreateOrderModal from "../../components/createOrderModal";
import CreateOrderButton from "../../components/createOrderButton";

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
                <CreateOrderButton setCreateModelOpen={setCreateModelOpen} />
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

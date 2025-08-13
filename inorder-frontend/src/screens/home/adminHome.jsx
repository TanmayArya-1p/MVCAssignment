import { useEffect, useState } from "react";
import Spinner from "@/components/spinner";
import { getAllOrders, getMyOrders } from "@/api/orders";
import OrderBook from "@/components/orderBook";
import CreateOrderModal from "@/components/createOrderModal";
import UnpaidBills from "@/components/unpaidBills";
import { useNavigate } from "react-router-dom";
import NavigateOrder from "@/components/navigateOrder";
import CreateOrderButton from "@/components/createOrderButton";

export default function AdminHomeScreen() {
    const navigate = useNavigate();
    const [ordersLoading, setOrdersLoading] = useState(true);
    const [orders, setOrders] = useState([]);
    const [createModelOpen, setCreateModelOpen] = useState(false);
    const [itemOrders, setItemOrders] = useState({});
    
    useEffect(() => {
        const fetchOrders = async () => {
            try {
                setOrdersLoading(true);
                const response = await getAllOrders(null, 0);
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
                        <CreateOrderButton setCreateModelOpen={setCreateModelOpen} />
                    </div>
                    {ordersLoading && <Spinner />}
                    </div>
                    <OrderBook orders={orders.filter(order => order.status !== "paid")} setOrders={setOrders} loading={ordersLoading} admin/>

                </div>

                <div className="flex flex-col gap-3">
                    <NavigateOrder />
                    <div className="mt-10">
                        <h2 className="text-3xl ubuntu-bold mb-2">Unpaid Bills</h2>
                        <UnpaidBills orders={orders} />

                    </div>


                </div>
            </div>
            <div className="flex flex-col">
                <div className="text-3xl ubuntu-bold mb-2">Previous Orders</div>
                <OrderBook noFilter orders={orders.filter(order => order.status === "paid")} setOrders={setOrders} loading={ordersLoading} admin />
            </div>
            <CreateOrderModal isOpen={createModelOpen} setIsOpen={setCreateModelOpen} />
            
        </div>

    </>
}

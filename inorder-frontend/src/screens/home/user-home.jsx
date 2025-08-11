import { useEffect, useState } from "react";
import axios from "axios";
import OrderCard from "../../components/orderCard";
import Spinner from "../../components/spinner";
import { getMyOrders } from "../../api/orders";
import OrderBook from "../../components/orderBook";
import ItemMenu from "../../components/itemMenu";

export default function UserHomeScreen() {
    const [ordersLoading, setOrdersLoading] = useState(true);
    const [orders, setOrders] = useState([]);

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
        <div className="flex flex-col mt-7 p-5">
            <div>
                <div className="text-3xl ubuntu-bold flex flex-row gap-5">Your Orders
                <button className="flex flex-row gap-2 text-lg items-center ubuntu-regular text-black">
                    <svg xmlns="http://www.w3.org/2000/svg" style={{width:20,height:20}} viewBox="0 0 640 640"><path d="M352 128C352 110.3 337.7 96 320 96C302.3 96 288 110.3 288 128L288 288L128 288C110.3 288 96 302.3 96 320C96 337.7 110.3 352 128 352L288 352L288 512C288 529.7 302.3 544 320 544C337.7 544 352 529.7 352 512L352 352L512 352C529.7 352 544 337.7 544 320C544 302.3 529.7 288 512 288L352 288L352 128z"/></svg>
                    <div>Create Order</div>
                </button>
            </div>
            {ordersLoading && <Spinner />}
            </div>
            <OrderBook orders={orders} loading={ordersLoading}/>

            {/* <button className="flex flex-row gap-3 mt-5 w-fit p-5">
                <svg style={{width:38,height:38}} xmlns="http://www.w3.org/2000/svg" viewBox="0 0 640 640"><path d="M127.9 78.4C127.1 70.2 120.2 64 112 64C103.8 64 96.9 70.2 96 78.3L81.9 213.7C80.6 219.7 80 225.8 80 231.9C80 277.8 115.1 315.5 160 319.6L160 544C160 561.7 174.3 576 192 576C209.7 576 224 561.7 224 544L224 319.6C268.9 315.5 304 277.8 304 231.9C304 225.8 303.4 219.7 302.1 213.7L287.9 78.3C287.1 70.2 280.2 64 272 64C263.8 64 256.9 70.2 256.1 78.4L242.5 213.9C241.9 219.6 237.1 224 231.4 224C225.6 224 220.8 219.6 220.2 213.8L207.9 78.6C207.2 70.3 200.3 64 192 64C183.7 64 176.8 70.3 176.1 78.6L163.8 213.8C163.3 219.6 158.4 224 152.6 224C146.8 224 142 219.6 141.5 213.9L127.9 78.4zM512 64C496 64 384 96 384 240L384 352C384 387.3 412.7 416 448 416L480 416L480 544C480 561.7 494.3 576 512 576C529.7 576 544 561.7 544 544L544 96C544 78.3 529.7 64 512 64z"/></svg>
                <div className="ubuntu-regular text-3xl">
                    Item Menu
                </div>
                <svg style={{width:38,height:38}} xmlns="http://www.w3.org/2000/svg" viewBox="0 0 640 640"><path d="M471.1 297.4C483.6 309.9 483.6 330.2 471.1 342.7L279.1 534.7C266.6 547.2 246.3 547.2 233.8 534.7C221.3 522.2 221.3 501.9 233.8 489.4L403.2 320L233.9 150.6C221.4 138.1 221.4 117.8 233.9 105.3C246.4 92.8 266.7 92.8 279.2 105.3L471.2 297.3z"/></svg>
            </button> */}

            <ItemMenu itemOrders={itemOrders} />

        </div>

    </>
}

import { deleteOrder } from "../api/orders";
import DeleteIcon from "../icons/deleteIcon";
import RightChevron from "../icons/rightChevron";
import {orderColourMap} from "../utils/const";
import { Toaster,toast } from "react-hot-toast";


export default function OrderCard({order,setOrders,admin}) {

    const deleteOrderHandler = async () => {
        if(!window.confirm("Are you sure you want to delete this order? This action cannot be undone.")) {
            return;
        }
        try {
            await deleteOrder(order.id)
            toast.success("Order deleted successfully");
            setOrders(prevOrders => prevOrders.filter(o => o.id !== order.id));
        } catch (error) {
            toast.error("Failed to delete order");
        }
    }



    return <div className="flex flex-row gap-1"><a href={`/order/${order.id}`} className="order-card">
        <div>
            <h2 className="text-lg font-bold ubuntu-bold">Order #{order.id} ( <span className={`text-${orderColourMap[order.status]}`}>{order.status}</span>  )</h2>
            <p className="mt-2 text-gray-500">Created At {new Date(order.issued_at).toLocaleString()}</p>
        </div>
        <RightChevron className="size-6 text-gray-500" />
    </a>
    {admin && <button className="delete-button" onClick={deleteOrderHandler}>
        <DeleteIcon className="size-6" />
    </button> }
    </div>
}
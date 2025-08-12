import { deleteOrder } from "../api/orders";
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
            console.log("error deleting order:", error);
            toast.error("Failed to delete order");
        }
    }



    return <div className="flex flex-row gap-1"><a href={`/order/${order.id}`} className="order-card">
        <Toaster></Toaster>
        <div>
            <h2 className="text-lg font-bold ubuntu-bold">Order #{order.id} ( <span className={`text-${orderColourMap[order.status]}`}>{order.status}</span>  )</h2>
            <p className="mt-2 text-gray-500">Created At {new Date(order.issued_at).toLocaleString()}</p>
        </div>
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth="1.5" stroke="currentColor" className="size-6">
            <path strokeLinecap="round" strokeLinejoin="round" d="m8.25 4.5 7.5 7.5-7.5 7.5" />
        </svg>
    </a>
    {admin && <button className="delete-button" onClick={deleteOrderHandler}>
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth="1.5" stroke="currentColor" className="size-6">
            <path strokeLinecap="round" strokeLinejoin="round" d="m14.74 9-.346 9m-4.788 0L9.26 9m9.968-3.21c.342.052.682.107 1.022.166m-1.022-.165L18.16 19.673a2.25 2.25 0 0 1-2.244 2.077H8.084a2.25 2.25 0 0 1-2.244-2.077L4.772 5.79m14.456 0a48.108 48.108 0 0 0-3.478-.397m-12 .562c.34-.059.68-.114 1.022-.165m0 0a48.11 48.11 0 0 1 3.478-.397m7.5 0v-.916c0-1.18-.91-2.164-2.09-2.201a51.964 51.964 0 0 0-3.32 0c-1.18.037-2.09 1.022-2.09 2.201v.916m7.5 0a48.667 48.667 0 0 0-7.5 0" />
        </svg>
    </button> }
    </div>
}
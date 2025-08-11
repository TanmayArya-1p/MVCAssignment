import {orderColourMap} from "../utils/const";

export default function OrderCard({order}) {
    return <a href={`/order/${order.id}`} className="order-card">
        <div>
            <h2 className="text-lg font-bold ubuntu-bold">Order {order.id} ( <span className={`text-${orderColourMap[order.status]}`}>{order.status}</span>  )</h2>
            <p className="mt-2 text-gray-500">Created At {new Date(order.issued_at).toLocaleString()}</p>
        </div>
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth="1.5" stroke="currentColor" className="size-6">
            <path stroke-linecap="round" stroke-linejoin="round" d="m8.25 4.5 7.5 7.5-7.5 7.5" />
        </svg>
    </a>
}
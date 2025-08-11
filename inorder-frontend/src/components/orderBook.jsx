import { ORDERS_PER_PAGE } from "../config";
import { Pagination } from '@mui/material';
import { useEffect, useState } from "react";
import OrderCard from "./orderCard";

export default function OrderBook({orders,loading}) {

    const [displayedOrders, setDisplayedOrders] = useState(orders.slice(0, ORDERS_PER_PAGE));
    const [ordersPage, setOrdersPage] = useState(1);
    const [displayedOrdersLength, setDisplayedOrdersLength] = useState();
    const [tags,setTags] = useState({
        "pending": false,
        "preparing": false,
        "served": false,
        "billed": false,
        "paid": false,
    });

    useEffect(() => {
        setDisplayedOrdersLength(orders.length);
    },[orders])

    useEffect(() => {
        console.log(tags)
        let selected = Object.keys(tags).filter(tag => tags[tag]);
        let temp = []
        if(selected.length === 0) {
            temp = orders
        } else {
            temp = orders.filter(order => selected.includes(order.status))
        }
        setOrdersPage(1);
        setDisplayedOrders(temp.slice(0, ORDERS_PER_PAGE));
        setDisplayedOrdersLength(temp.length);

    }, [tags]);

    useEffect(() => {
        setDisplayedOrders(orders.slice((ordersPage - 1) * ORDERS_PER_PAGE, ordersPage * ORDERS_PER_PAGE));
    }, [orders, ordersPage]);

    if(loading) {
        return <></>
    }
    return <>
        <div id='tag-container' className="mt-5 flex flex-row gap-3 items-center">
            <div className="ubuntu-bold text-md">
                Filters:
            </div>
            {Object.keys(tags).filter(tag => tags[tag]).map(tag => (
                <div key={tag} className='tag tag-selected' onClick={() => setTags((prevTags) => ({
                    ...prevTags,
                    [tag]: !prevTags[tag],
                }))}>
                    {tag}
                </div>
            ))}
            {Object.keys(tags).filter(tag => !tags[tag]).map(tag => (
                <div key={tag} className='tag' onClick={() => setTags((prevTags) => ({
                    ...prevTags,
                    [tag]: !prevTags[tag],
                }))}>
                    {tag}
                </div>
            ))}

        </div>

        <div id="orders-container" className="flex flex-row mt-2 flex-wrap gap-3 items-center">
            {displayedOrders.map(order => <OrderCard key={order.id} order={order} />)}
        </div>
        <div className="flex flex-row gap-3 mt-2">
            <Pagination color="standard" count={Math.ceil(displayedOrdersLength / ORDERS_PER_PAGE)} variant="outlined" shape="rounded" page={ordersPage} onChange={(event, value) => setOrdersPage(value)} />
        </div>
    </>
}
import { ORDERS_PER_PAGE } from "../config";
import { Pagination } from '@mui/material';
import { useEffect, useState } from "react";
import OrderCard from "./orderCard";

export default function OrderBook({noFilter, setOrders,orders, loading, admin}) {

    const [displayedOrders, setDisplayedOrders] = useState(orders.slice(0, ORDERS_PER_PAGE));
    const [filteredOrders, setFilteredOrders] = useState(orders);
    const [ordersPage, setOrdersPage] = useState(1);
    const [tags,setTags] = useState({
        "pending": false,
        "preparing": false,
        "served": false,
        "billed": false,
    });

    useEffect(() => {
        if(!admin) setTags(t=> ({...t,"paid": false}));
    },[])

    useEffect(() => {
        setFilteredOrders(orders);
    },[orders])

    useEffect(() => {
        setDisplayedOrders(filteredOrders.slice((ordersPage - 1) * ORDERS_PER_PAGE, ordersPage * ORDERS_PER_PAGE));
        setOrdersPage(1);
    },[filteredOrders])

    useEffect(() => {
        setDisplayedOrders(filteredOrders.slice((ordersPage - 1) * ORDERS_PER_PAGE, ordersPage * ORDERS_PER_PAGE));
    }, [filteredOrders, ordersPage]);

    useEffect(() => {
        let selected = Object.keys(tags).filter(tag => tags[tag]);
        let temp = []
        if(selected.length === 0) {
            temp = orders
        } else {
            temp = orders.filter(order => selected.includes(order.status))
        }
        setFilteredOrders(temp);
    }, [tags]);


    if(loading) {
        return <></>
    }
    return <>
        {!noFilter && 
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
        }

        <div id="orders-container" className="flex flex-row mt-2 flex-wrap gap-3 items-center">
            {displayedOrders.map(order => <OrderCard key={order.id} order={order} admin={admin} setOrders={setOrders}/>)}
        </div>
        <div className="flex flex-row gap-3 mt-2">
            <Pagination color="standard" className="bg-white rounded-sm border-2 p-2" count={Math.ceil(filteredOrders.length / ORDERS_PER_PAGE)} variant="outlined" shape="rounded" page={ordersPage} onChange={(event, value) => setOrdersPage(value)} />
        </div>
    </>
}
import { useState, useEffect } from 'react';
import { getAllOrderedItems } from '../api/items';
import { orderItemStatusPriority, orderColourMap, bumpItemStatusMap } from '../utils/const';
import { bumpOrderItemStatus } from '../api/items';
import {toast, Toaster} from 'react-hot-toast';

export function ItemQueue() {

    const [orderedItems, setOrderedItems] = useState([]);
    
    useEffect(() => {
        async function fetchItems() {
            try {
                let response = await getAllOrderedItems();
                response = response.filter(item => item.status !== "served");
                response.sort((a, b) => orderItemStatusPriority[a.status] - orderItemStatusPriority[b.status]);

                setOrderedItems(response);
            } catch(error) {
                console.error("error fetching ordered items:", error);
                setOrderedItems([]);
            } 

        }
        fetchItems();
    }, [setOrderedItems]);

    const bumpItemStatus = async (itemId) => {
        try {
            await bumpOrderItemStatus(itemId);
            setOrderedItems(prevItems => {
                let updatedItems = prevItems.map(item => 
                    item.id === itemId ? { ...item, status: bumpItemStatusMap[item.status] } : item
                );
                updatedItems= updatedItems.filter(item => item.status !== "served");
                updatedItems.sort((a, b) => orderItemStatusPriority[a.status] - orderItemStatusPriority[b.status]);
                return [...updatedItems];
            });
            toast.success("Item status bumped successfully.");
        } catch (error) {
            toast.error("Failed to bump item status.");
            console.error("error bumping item status:", error);
        }

    }


    return <div id="item-queue-container" className="relative overflow-x-auto max-w-fit p-3 bg-white rounded-sm shadow-md border-2" style={{ maxHeight: "40rem" }}>
            <table>
                <thead>
                    <tr className='text-lg'>
                        <th className="ubuntu-bold px-3 text-left">Item</th>
                        <th className="ubuntu-bold px-3">Quantity</th>
                        <th className="ubuntu-bold px-3">Instructions</th>
                        <th className="ubuntu-bold px-3">Ordered At</th>
                        <th className="ubuntu-bold px-3">Order</th>
                        <th className="ubuntu-bold px-3">Status</th>
                        <th></th>
                    </tr>
                </thead>
                <tbody>
                    {orderedItems.map((item) => (
                        <tr className="ubuntu-regular text-left" key={item.id}>
                            <td className="px-3 ubuntu-bold">{item.name}</td>
                            <td className="px-3 text-center">{item.quantity}</td>
                            <td className="px-3 truncate">{item.instructions || "No instructions"}</td>
                            <td className="px-3">{new Date(item.issued_at).toLocaleString()}</td>
                            <td className="px-3">
                                <a className="text-bold order-link" href={`/order/${item.order_id}`}>Order #{item.order_id}</a>
                            </td>
                            <td className="px-3 text-lg font-bold ubuntu-bold text-center">
                                <span className={`text-${orderColourMap ? orderColourMap[item.status] : ''}`}>{item.status}</span>
                            </td>
                            <td className="px-3 py-1">
                                <button onClick={() => bumpItemStatus(item.id)} className='!py-0'>Bump Status</button>
                            </td>
                        </tr>
                    ))}

                </tbody>
            </table>
            { orderedItems.length === 0 && (
                <div className="ubuntu-bold text-center text-xl mt-5">No items in queue</div>
            )}
        </div>



}
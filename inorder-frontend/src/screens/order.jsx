import { useParams,useSearchParams } from 'react-router-dom';
import Navbar from '../components/navbar';
import Spinner from '../components/spinner';
import {  useEffect,useState } from 'react';
import { addItemToOrder, getOrder } from '../api/orders';
import { orderColourMap } from '../utils/const';
import Bill from '../components/bill';
import useAuthStore from '../stores/authStore';
import Modal from 'react-modal';
import ItemMenu from '../components/itemMenu';
import { modalStyle } from '../utils/const';
import { Toaster,toast } from 'react-hot-toast';


export default function OrderScreen() {
    const {orderid} = useParams();

    const [searchParams] = useSearchParams();
    const add = searchParams.get("add")=== "true";

    const [loading, setLoading] = useState(true);
    const [order, setOrder] = useState(null);
    const {role} = useAuthStore.getState();

    const [addItemsModalIsOpen, setAddItemsModalIsOpen] = useState(add);
    const [itemOrders, setItemOrders] = useState([]);
    const [itemInstructions, setItemInstructions] = useState({});
    const [addedItemPrice, setAddedItemPrice] = useState(0);
    const addItemsHandler = async () => {
        if(!window.confirm(`Are you sure you want to add these items(₹${addedItemPrice}) to the order?`)) {
            return;
        }
        try {
            Object.keys(itemOrders).forEach(async (itemId) => {
                if(itemOrders[itemId] > 0) {
                    await addItemToOrder(orderid, itemId, itemOrders[itemId], itemInstructions[itemId] || "");
                }
            });
        } catch (error) {
            console.error("error adding items to order:", error);
            toast.error("Failed to add items to order.");
            return;
        }

        setAddItemsModalIsOpen(false);
        setItemOrders({});
        setItemInstructions({});
        toast.success("Items added to order successfully");
        setTimeout(() => {
            window.location.href= `/order/${orderid}`;
        }, 1300);
    }
    

    const [billLoading, setBillLoading] = useState(true);

    useEffect(()=> {
        const fetchOrder = async () => {
            try {
                const response = await getOrder(orderid);
                setOrder(response);
            } catch (error) {
                console.error("error getting order", error);
                window.location.href = "/notfound"
            } finally {
                setLoading(false);
            }
        }
        fetchOrder();
    },[orderid])


    if(loading) {
        return <div className='mt-10'><Spinner /></div>
    }


    return <div className="h-screen w-screen flex flex-col">
            <Toaster />
            <Navbar/>
            <div className='mt-5 p-5'>
                <div>
                    <div className='ubuntu-bold text-4xl'>
                        Order #{order.id} ( <span className={'text-'+orderColourMap[order.status]}>{order.status}</span> )
                    </div>
                    <div className='ubuntu-regular text-lg mt-2'>
                        Ordered at Table {order.table_no} at {new Date(order.issued_at).toLocaleString()}
                    </div>
                </div>
                <div>
                    <div className='ubuntu-bold text-4xl mt-9 flex flex-row gap-3 items-center'>
                        Current Bill
                    </div>
                    <Bill order={order} setBillLoading={setBillLoading} billLoading={billLoading} role={role} />
                </div>
                {(order.status !== "paid" && order.status !=="billed") ?
                <button onClick={() => setAddItemsModalIsOpen(true)} className='mt-6 flex flex-row items-center gap-2 text-xl w-fit'>
                    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="size-6">
                    <path strokeLinecap="round" strokeLinejoin="round" d="M12 4.5v15m7.5-7.5h-15" />
                    </svg>

                    Add Items to Order
                </button>       
                : null}
            </div>
     
            <Modal
                isOpen={addItemsModalIsOpen}
                onRequestClose={() => setAddItemsModalIsOpen(false)}
                style={modalStyle}
                ariaHideApp={false}
                contentLabel="Add Items"
            >
                <ItemMenu itemOrders={itemOrders} setItemOrders={setItemOrders} itemInstructions={itemInstructions} setItemInstructions={setItemInstructions} setAddedItemPrice={setAddedItemPrice} pageSize={4}/>
                <div className='flex flex-row gap-5 justify-end'>
                    <button onClick={() => addItemsHandler()}>Add Items {"( +₹"+addedItemPrice+ " )"}</button>
                    <button onClick={() => setAddItemsModalIsOpen(false)}>Cancel</button>
                </div>
            </Modal>
    </div>
}
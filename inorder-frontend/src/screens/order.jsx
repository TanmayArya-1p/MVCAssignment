import { useParams,useSearchParams } from 'react-router-dom';
import Navbar from '../components/navbar';
import Spinner from '../components/spinner';
import {  useEffect,useState } from 'react';
import { addItemToOrder, getOrder } from '../api/orders';
import { orderColourMap, roles } from '../utils/const';
import Bill from '../components/bill';
import useAuthStore from '../stores/authStore';
import Modal from 'react-modal';
import ItemMenu from '../components/itemMenu';
import { modalStyle } from '../utils/const';
import { Toaster,toast } from 'react-hot-toast';
import { deleteOrder } from '../api/orders';
import VerifySignedIn from "../utils/verify";
import { useNavigate } from 'react-router-dom';

export default function OrderScreen() {
    const {orderid} = useParams();
    const navigate = useNavigate();
    const [searchParams] = useSearchParams();
    const add = searchParams.get("add")=== "true";
    useEffect(() => {VerifySignedIn()}, [])


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
            window.location.pathname = `/order/${orderid}`;
        }, 1000);
    }
    

    const [billLoading, setBillLoading] = useState(true);

    useEffect(()=> {
        const fetchOrder = async () => {
            try {
                const response = await getOrder(orderid);
                setOrder(response);
            } catch (error) {
                console.error("error getting order", error);
                navigate("/notfound");
            } finally {
                setLoading(false);
            }
        }
        fetchOrder();
    },[orderid])

    const deleteOrderHandler = async () => {
        if(!window.confirm("Are you sure you want to delete this order? This action cannot be undone.")) {
            return;
        }
        try {
            await deleteOrder(order.id)
            toast.success("Order deleted successfully");
            setTimeout(() => {
                navigate("/home");
            }, 1000);

        } catch (error) {
            toast.error("Failed to delete order");
        }
    }

    if(loading) {
        return <div className='mt-10'><Spinner /></div>
    }


    return <div className="h-screen w-screen flex flex-col">
            <title>Order #{order.id} - InOrder</title>
            <Toaster />
            <Navbar/>
            <div className='mt-5 p-5'>
                <div>
                    <div className='ubuntu-bold text-4xl flex flex-row gap-3'>
                        Order #{order.id} ( <span className={'text-'+orderColourMap[order.status]}>{order.status}</span> )
                        {role===roles.ADMIN &&
                            <button onClick={deleteOrderHandler}>
                                <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="size-6">
                                    <path strokeLinecap="round" strokeLinejoin="round" d="m14.74 9-.346 9m-4.788 0L9.26 9m9.968-3.21c.342.052.682.107 1.022.166m-1.022-.165L18.16 19.673a2.25 2.25 0 0 1-2.244 2.077H8.084a2.25 2.25 0 0 1-2.244-2.077L4.772 5.79m14.456 0a48.108 48.108 0 0 0-3.478-.397m-12 .562c.34-.059.68-.114 1.022-.165m0 0a48.11 48.11 0 0 1 3.478-.397m7.5 0v-.916c0-1.18-.91-2.164-2.09-2.201a51.964 51.964 0 0 0-3.32 0c-1.18.037-2.09 1.022-2.09 2.201v.916m7.5 0a48.667 48.667 0 0 0-7.5 0" />
                                </svg>
                            </button>
                        }
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
                    <button onClick={() => addItemsHandler()}>Add Items {"( +₹"+addedItemPrice.toFixed(2)+ " )"}</button>
                    <button onClick={() => setAddItemsModalIsOpen(false)}>Cancel</button>
                </div>
            </Modal>
    </div>
}
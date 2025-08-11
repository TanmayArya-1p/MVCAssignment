import { useParams } from 'react-router-dom';
import Navbar from '../components/navbar';
import Spinner from '../components/spinner';
import {  useEffect,useState } from 'react';
import { getOrder } from '../api/orders';
import { orderColourMap } from '../utils/const';
import Bill from '../components/bill';
import useAuthStore from '../stores/authStore';

export default function OrderScreen() {
    const {orderid} = useParams();
    const [loading, setLoading] = useState(true);
    const [order, setOrder] = useState(null);
    const {role} = useAuthStore.getState();

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
            </div>
    </div>
}
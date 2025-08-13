import Modal from 'react-modal';
import {modalStyle} from '@/utils/const';
import { useState } from 'react';
import { createOrder } from '@/api/orders';
import { toast } from 'react-hot-toast';
import { useNavigate } from 'react-router-dom';

export default function CreateOrderModal({ isOpen, setIsOpen }) {

    const [tableNo, setTableNo] = useState("");
    const navigate = useNavigate();

    const createOrderHandler = () => {
        async function create() {
            try {
                let resp = await createOrder(tableNo);
                toast.success("Order Successfully Created");
                setTimeout(() => {
                    setIsOpen(false);
                    navigate(`/order/${resp.id}?add=true`);
                }, 1000);
            } catch (error) {
                toast.error("Error Creating Order");
            }
        }
        create();
    }

    return (<>
        <Modal
            isOpen={isOpen}
            onRequestClose={() => setIsOpen(false)}
            style={modalStyle}
            contentLabel="Create Order Modal"
      >
        <div className='flex align-center flex-col'>
            <div className='ubuntu-bold text-lg text-center'>Create Order</div>
            <div>
                <div className='flex flex-row gap-5 mt-5 align-center justify-center'>
                    <input type="number" placeholder='Table No' value={tableNo} onChange={(e) => setTableNo(e.target.value>=0 ? e.target.value : 1)} />
                    <button onClick={createOrderHandler}>Create</button>
                    <button onClick={() => setIsOpen(false)}>Close</button>
                </div>

            </div>
        </div>

      </Modal>
    </>);
}
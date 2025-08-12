import Modal from 'react-modal';
import {modalStyle} from '../utils/const';
import { useState } from 'react';
import { createOrder } from '../api/orders';
import { toast, Toaster } from 'react-hot-toast';

export default function CreateOrderModal({ isOpen, setIsOpen }) {

    const [tableNo, setTableNo] = useState(1);

    const createOrderHandler = () => {
        async function create() {
            try {
                let resp = await createOrder(tableNo);
                toast.success("Order Successfully Created");
                setTimeout(() => {
                    setIsOpen(false);
                    window.location.href = `/order/${resp.id}?add=true`;
                }, 1300);
            } catch (error) {
                toast.error("Error Creating Order");
            }
        }
        create();
    }

    return (<>
        <Toaster />
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
                    <input type="number" placeholder='Table No' value={tableNo} className='bg-white p-2 border-2' onChange={(e) => setTableNo(e.target.value>=0 ? e.target.value : 1)} />
                    <button onClick={createOrderHandler}>Create</button>
                    <button onClick={() => setIsOpen(false)}>Close</button>
                </div>

            </div>
        </div>

      </Modal>
    </>);
}
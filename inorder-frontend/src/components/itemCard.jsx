import { Toaster,toast } from "react-hot-toast";
import { deleteItem } from "../api/items";
import { API_URL } from "../config"


export default function ItemCard({admin,setItems, item,setItemOrders,itemOrders,itemInstructions,setItemInstructions, setAddedItemPrice}) {


    const deleteItemHandler = async () => {
        if(!window.confirm("Are you sure you want to delete this item?")) {
            return;
        }
        try {
            await deleteItem(item.id)
            setItems((prevItems) => prevItems.filter((i) => i.id !== item.id));
            toast.success("Item deleted successfully");
        } catch (error) {
            console.error("error deleting item:", error);
            toast.error("Failed to delete item");
        }
    }


    return (<div className="flex flex-col gap-2">
        <Toaster></Toaster>
        <div className="item-card">
                <img src={API_URL+"/public"+item.image} style={{width: '9rem', maxHeight: '10rem'}} />
                <div className="w-full">
                    <h2 className="text-lg font-bold ubuntu-bold truncate">{item.name}</h2>
                    <p className="mt-2 text-gray-500 truncate">{item.description}</p>
                    <p className="mt-2 text-gray-800">Price: ₹{item.price}</p>
                    <div id="tags" className="flex overflow-x-auto flex-row gap-2 mt-1">
                        {item.tags.map((tag, ind) => (
                            <span key={ind} className="tag-no-hover">{tag}</span>
                        ))}
                    </div>

                </div>
                {itemOrders &&
                    <div className="relative flex items-center text-xl">
                        <button className="ubuntu-bold !p-0.5 w-10 disabled:text-gray-400" onClick={() => {
                            setItemOrders((prevOrders) => ({
                                ...prevOrders,
                                [item.id]: (prevOrders[item.id] || 0) - 1,
                            }));
                            setAddedItemPrice((prevPrice) => prevPrice - item.price);
                        }} disabled={!itemOrders[item.id] || itemOrders[item.id] <= 0}>
                            -
                        </button>
                        <label type="text" className="px-4 ubuntu-bold"  placeholder="0" value="12">
                            {itemOrders[item.id] || 0}
                        </label>
                        <button className="ubuntu-bold !p-0.5 w-10" onClick={() => {
                            setItemOrders((prevOrders) => ({
                                ...prevOrders,
                                [item.id]: (prevOrders[item.id] || 0) + 1,
                            }));
                            setAddedItemPrice((prevPrice) => prevPrice + item.price);

                        }}>
                            +
                        </button>
                    </div>
                }
                {(itemOrders && itemOrders[item.id]) ? (
                    <input type="text" value={itemInstructions[item.id] || ""} onChange={(e) => {
                        setItemInstructions((prevInstructions) => ({
                            ...prevInstructions,
                            [item.id]: e.target.value
                        }));
                    }} className="bg-white border-2 w-45 rounded-sm mt-3 p-1 px-2" placeholder="Instructions" />
                ) : null}
        </div>
        {(admin) && (
            <div className="flex flex-row gap-2">
                <button className="justify-center flex delete-button w-full" onClick={deleteItemHandler}>
                    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth="1.5" stroke="currentColor" className="size-6">
                        <path strokeLinecap="round" strokeLinejoin="round" d="m14.74 9-.346 9m-4.788 0L9.26 9m9.968-3.21c.342.052.682.107 1.022.166m-1.022-.165L18.16 19.673a2.25 2.25 0 0 1-2.244 2.077H8.084a2.25 2.25 0 0 1-2.244-2.077L4.772 5.79m14.456 0a48.108 48.108 0 0 0-3.478-.397m-12 .562c.34-.059.68-.114 1.022-.165m0 0a48.11 48.11 0 0 1 3.478-.397m7.5 0v-.916c0-1.18-.91-2.164-2.09-2.201a51.964 51.964 0 0 0-3.32 0c-1.18.037-2.09 1.022-2.09 2.201v.916m7.5 0a48.667 48.667 0 0 0-7.5 0" />
                    </svg>
                </button>
                <button className="justify-center flex w-full gap-2" onClick={()=> window.location.href=`/items/update?id=${item.id}&name=${item.name}&price=${item.price}&description=${item.description}&image=${item.image}&tags=${item.tags.join(",")}`}>
                    Edit
                    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="size-6">
                    <path strokeLinecap="round" strokeLinejoin="round" d="m16.862 4.487 1.687-1.688a1.875 1.875 0 1 1 2.652 2.652L10.582 16.07a4.5 4.5 0 0 1-1.897 1.13L6 18l.8-2.685a4.5 4.5 0 0 1 1.13-1.897l8.932-8.931Zm0 0L19.5 7.125M18 14v4.75A2.25 2.25 0 0 1 15.75 21H5.25A2.25 2.25 0 0 1 3 18.75V8.25A2.25 2.25 0 0 1 5.25 6H10" />
                    </svg>
                </button>
            </div>


        )}
        </div>
    )
}
import { Toaster,toast } from "react-hot-toast";
import { deleteItem } from "@/api/items";
import { API_URL } from "@/config"
import { useNavigate } from "react-router-dom";
import DeleteIcon from "@/icons/deleteIcon";
import EditIcon from "@/icons/editIcon";

export default function ItemCard({admin,setItems, item,setItemOrders,itemOrders,itemInstructions,setItemInstructions, setAddedItemPrice}) {
    const navigate = useNavigate();

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
                    }} className="w-45 mt-3 !p-0.5 !px-2" placeholder="Instructions" />
                ) : null}
        </div>
        {(admin) && (
            <div className="flex flex-row gap-2">
                <button className="justify-center flex delete-button w-full" onClick={deleteItemHandler}>
                    <DeleteIcon className="size-6" />
                </button>
                <button className="justify-center flex w-full gap-2" 
                        onClick={()=> navigate(`/items/update?id=${item.id}&name=${item.name}&price=${item.price}&description=${item.description}&image=${item.image}&tags=${item.tags.join(",")}`)}>
                    Edit
                    <EditIcon className="size-6" />
                </button>
            </div>


        )}
        </div>
    )
}
import { API_URL } from "../config"


export default function ItemCard({item,setItemOrders,itemOrders}) {
    return (
        <div className="item-card">
                <img src={API_URL+"/public"+item.image} style={{width: '9rem', maxHeight: '10rem'}} />
                <div className="w-full">
                    <h2 className="text-lg font-bold ubuntu-bold">{item.name}</h2>
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
                        }}>
                            +
                        </button>
                    </div>
                }
            
        </div>
    )
}
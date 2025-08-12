import { ITEMS_PER_PAGE } from "../config";
import { Pagination } from '@mui/material';
import { useEffect, useState } from "react";
import ItemCard from "./itemCard";
import { getAllItems, getItemsOfTags } from "../api/items";
import { getAllTags } from "../api/tags";

export default function ItemMenu({itemOrders,setItemOrders,itemInstructions,setItemInstructions,setAddedItemPrice,pageSize}) {

    const [items, setItems] = useState([]);
    const [loading,setLoading] = useState(true);
    const [tags,setTags] = useState({});

    const [displayedItems, setDisplayedItems] = useState([]); //subset of indexed items
    const [filteredItems, setFilteredItems] = useState([]); //subset of items
    const [indexedItems, setIndexedItems] = useState([]); //subset of filtered items


    const [itemsPage, setItemsPage] = useState(1);

    const [query, setQuery] = useState("");
    
    const handleSearch = (e) => {
        // if (e.key == "Enter") {
        //     triggerSearch();
        //     return;
        // }
        setQuery(e.target.value);
    };

    const triggerSearch = () => {
        if (query.trim() === "") {
            setIndexedItems(filteredItems);
            return;
        }
        let filtered = filteredItems.filter(item => item.name.toLowerCase().includes(query.toLowerCase()));
        setIndexedItems(filtered);
        setItemsPage(1);
    }

    
    useEffect(() => {
        triggerSearch();
    
    }, [query]);


    useEffect(() => {
        const fetchItems = async () => {
            try {
                const response = await getAllItems();
                setItems(response);
            } catch (error) {
                console.error("error getting items:", error);
            }

            try {
                const response = await getAllTags();
                const tagsObj = {};
                response.forEach(tag => {
                    tagsObj[tag] = false;
                });
                setTags(tagsObj);
            } catch (error) {
                console.error("error getting tags:", error);
            } 
            setLoading(false);     
        };

        fetchItems();
    }, []);



    useEffect(() => {
        setFilteredItems(items);
        setIndexedItems(filteredItems);
        setDisplayedItems(indexedItems.slice(0, pageSize || ITEMS_PER_PAGE));
    },[items])

    useEffect(() => {
        setIndexedItems(filteredItems);
        setItemsPage(1);
    }, [filteredItems]);

    useEffect(() => {
        console.log(tags)
        let selected = Object.keys(tags).filter(tag => tags[tag]);
        if(selected.length === 0) {
            setItemsPage(1);
            setFilteredItems(items);
        } else {
            async function fetchItemsByTags() {
                setLoading(true);
                let temp = await getItemsOfTags(selected);
                setItemsPage(1);
                setFilteredItems(temp);
                setLoading(false);
            }
            fetchItemsByTags();
        }
        setQuery("");
    }, [tags, items]);

    useEffect(() => {
        setDisplayedItems(indexedItems.slice((itemsPage - 1) *  (pageSize || ITEMS_PER_PAGE), itemsPage * (pageSize || ITEMS_PER_PAGE)));
    }, [indexedItems, itemsPage]);

    if(loading) {
        return <></>
    }
    return <>
        <div id="tag-container" className="mt-5 flex flex-row flex-wrap gap-3 items-center">
            <div className="ubuntu-bold text-md flex items-center">
                Tags:
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

        <div id="query-container">
            <div className="flex flex-row gap-2 items-center mt-3 bg-white border-2 rounded-sm p-2 w-fit">
                <svg xmlns="http://www.w3.org/2000/svg"  fill="none" viewBox="0 0 24 24" strokeWidth={2} stroke="currentColor" className="size-6">
                    <path strokeLinecap="round" strokeLinejoin="round" d="m21 21-5.197-5.197m0 0A7.5 7.5 0 1 0 5.196 5.196a7.5 7.5 0 0 0 10.607 10.607Z" />
                </svg>
                <input type="text" id="search-input" value={query} onChange={handleSearch} placeholder="Search Items" className="focus:outline-0"/>
            </div>
        </div>

        <div id="items-container" className="mt-3 flex flex-row flex-wrap gap-3 items-center">
            {displayedItems.map(item => <ItemCard key={item.id} item={item} setItemOrders={setItemOrders} itemOrders={itemOrders} itemInstructions={itemInstructions} setItemInstructions={setItemInstructions} setAddedItemPrice={setAddedItemPrice} />)}
            {displayedItems.length === 0 && <div className="ubuntu-bold h-80 w-full text-center text-3xl flex justify-center items-center">No items found</div>}
         </div>
        <div className="flex flex-row gap-3 mt-12">
            <Pagination color="standard" count={Math.ceil(indexedItems.length / (pageSize || ITEMS_PER_PAGE))} variant="outlined" shape="rounded" page={itemsPage} onChange={(event, value) => setItemsPage(value)} />
        </div>
    </>
}
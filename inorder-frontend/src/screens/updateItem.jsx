import { useSearchParams } from "react-router-dom";
import Navbar from "../components/navbar";
import VerifySignedIn from "../utils/verify";
import useAuthStore from "../stores/authStore";
import { useEffect } from "react";
import { Toaster } from "react-hot-toast";
import { API_URL } from "../config";
import { updateItem, uploadImage } from "../api/items";
import {toast} from "react-hot-toast";


//TODO: REPLACE WINDOW.LOCATIONS STUFF

export default function UpdateItemScreen() {
    const {role} = useAuthStore.getState();
    useEffect(() => {
        if (role !== "admin") {
            window.location.href = "/home";
        }
    }, [role]);

    useEffect(() => {VerifySignedIn()}, [])


    const [searchParams] = useSearchParams();
    const itemId = searchParams.get("id");
    const itemName = searchParams.get("name");
    const itemPrice = searchParams.get("price");
    const itemDescription = searchParams.get("description");
    const itemImage = searchParams.get("image");
    const itemTags = searchParams.get("tags");


    const handleSubmit = async (e) => {
        e.preventDefault();
        const formData = new FormData(e.target);
        const image = formData.get("image");
        try {
            let resp = null
            if(image.name) {
                console.log(image)
                resp = await uploadImage(image);
            }
            const createResp = await updateItem({
                name: formData.get("name"),
                price: formData.get("price"),
                description: formData.get("description"),
                tags: formData.get("item-tags").split(",").map(tag => tag.trim()),
                image: resp ? resp.url : itemImage,
                itemId
            });
            console.log("updating",createResp);
            toast.success("Successfully updated item");
            window.location = "/items"
        } catch(error) {
            console.error("error updating item:", error);
            toast.error("Failed to update item");
        }

    }

    useEffect(() => {
        if (!itemId || !itemName || !itemPrice || !itemDescription || !itemImage || !itemTags) {
            window.location.href = ("/notfound");
        }
    }, []);

    return <div className="h-screen w-screen flex flex-col">
        <Navbar></Navbar>
        <Toaster />
        <div className="mt-10 p-5 flex flex-col flex-wrap gap-10 items-center justify-center">
            <div className="ubuntu-bold text-4xl">Update Item #{itemId}</div>
            <div className="bg-white rounded-sm border-2 p-3">
                <div className="flex flex-row gap-5">
                    <form onSubmit={handleSubmit} id="create-item-form" className="flex flex-col gap-3 ubuntu-regular bg-white border-2 p-3 rounded" style={{width:"30rem"}}>
                        <div className="flex flex-row justify-between">
                            <div>Name:</div>
                            <input type="text" id="name" name="name" className="border-2 rounded-sm p-1 px-2"required placeholder={itemName}/>
                        </div>
                        <div className="flex flex-row justify-between">
                            <div>Description:</div>
                            <input type="text" id="description" name="description" className="border-2 rounded-sm p-1 px-2" required placeholder={itemDescription}/>
                        </div>
                        <div className="flex flex-row justify-between">
                            <div>Price:</div>
                            <input type="text" id="price" name="price" className="border-2 rounded-sm p-1 px-2" required placeholder={itemPrice}/>
                        </div>
                        <div className="flex flex-row justify-between">
                            <div>Tags:</div>
                            <input type="text" id="item-tags" name="item-tags" className="border-2 rounded-sm p-1 px-2" required placeholder={itemTags}/>
                        </div>
                        <div className="flex flex-row justify-between">
                            <div>New Image:</div>
                            <input type="file" id="file" name="image" accept="image/*" className="file-input"/>
                        </div>
                        <button type="submit" className="sbutton ubuntu-bold">Update Item</button>
                    </form>
                    <div className="flex items-center justify-center p-5">
                        <img src={API_URL+"/public"+itemImage} alt="Item Image" className="w-80 max-h-80"/>
                    </div>
                </div>
            </div>
        </div>
    </div>


}
import { useSearchParams } from "react-router-dom";
import Navbar from "@/components/navbar";
import VerifySignedIn from "@/utils/verify";
import useAuthStore from "@/stores/authStore";
import { useEffect } from "react";
import { Toaster } from "react-hot-toast";
import { API_URL } from "@/config";
import { updateItem, uploadImage } from "@/api/items";
import {toast} from "react-hot-toast";
import { useNavigate } from "react-router-dom";
import { roles } from "@/utils/const";



export default function UpdateItemScreen() {
    const navigate = useNavigate();
    const {role} = useAuthStore.getState();
    useEffect(() => {
        if (role !== roles.ADMIN) {
            navigate("/home");
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
            toast.success("Successfully updated item");
            navigate("/items");
        } catch(error) {
            toast.error("Failed to update item. Try a different Item Name");
        }

    }

    useEffect(() => {
        if (!itemId || !itemName || !itemPrice || !itemDescription || !itemImage || !itemTags) {
            navigate("/notfound");
        }
    }, []);

    return <div className="h-screen w-screen flex flex-col">
        <Navbar/>
        <title>Update Item - InOrder</title>
        <Toaster />
        <div className="mt-10 p-5 flex flex-col flex-wrap gap-10 items-center justify-center">
            <div className="ubuntu-bold text-4xl">Update Item #{itemId}</div>
            <div className="bg-white rounded-sm border-2 p-3">
                <div className="flex flex-row gap-5">
                    <form onSubmit={handleSubmit} 
                        className="flex flex-col gap-3 ubuntu-regular bg-white border-2 p-3 rounded" 
                        style={{width:"30rem"}}
                    >
                        <div className="form-div">
                            <div>Name:</div>
                            <input type="text" id="name" name="name" required placeholder={itemName}/>
                        </div>
                        <div className="form-div">
                            <div>Description:</div>
                            <input type="text" id="description" name="description" required placeholder={itemDescription}/>
                        </div>
                        <div className="form-div">
                            <div>Price:</div>
                            <input type="text" id="price" name="price" required placeholder={itemPrice}/>
                        </div>
                        <div className="form-div">
                            <div>Tags:</div>
                            <input type="text" id="item-tags" name="item-tags" required placeholder={itemTags}/>
                        </div>
                        <div className="form-div">
                            <div>New Image:</div>
                            <input type="file" id="file" name="image" accept="image/*" className="file-input"/>
                        </div>
                        <button type="submit" className="ubuntu-bold">Update Item</button>
                    </form>
                    <div className="flex items-center justify-center p-5">
                        <img src={API_URL+"/public"+itemImage} alt="Item Image" className="w-80 max-h-80"/>
                    </div>
                </div>
            </div>
        </div>
    </div>


}
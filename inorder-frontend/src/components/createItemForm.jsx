
import { uploadImage, createItem } from "@/api/items";
import {toast,Toaster} from "react-hot-toast";
import { useNavigate } from "react-router-dom";

export default function CreateItemForm() {
    const navigate = useNavigate();

    const handleSubmit = async (e) => {
        e.preventDefault();
        const formData = new FormData(e.target);
        const image = formData.get("image");
        try {
            let resp = null
            if(image.name.length > 0) {
                resp = await uploadImage(image);
            }
            const createResp = await createItem({
                name: formData.get("name"),
                price: formData.get("price"),
                description: formData.get("description"),
                tags: formData.get("item-tags").split(",").map(tag => tag.trim()),
                image: resp ? resp.url : null,
            });
            navigate(0)
        } catch(error) {
            toast.error("Item with same name exists or another error occured");
        }

    }

    return <div id="create-item-container" className="flex flex-col gap-4">
            <form onSubmit={handleSubmit} id="create-item-form" className="flex flex-col gap-3 ubuntu-regular bg-white border-2 p-3 rounded" style={{width:"30rem"}}>
                <div className="flex flex-row justify-between">
                    <div>Name:</div>
                    <input type="text" id="name" name="name" required placeholder="Burger"/>
                </div>
                <div className="flex flex-row justify-between">
                    <div>Description:</div>
                    <input type="text" id="description" name="description" required placeholder="just a burger"/>
                </div>
                <div className="flex flex-row justify-between">
                    <div>Price:</div>
                    <input type="text" id="price" name="price" required placeholder="30"/>
                </div>
                <div className="flex flex-row justify-between">
                    <div>Tags:</div>
                    <input type="text" id="item-tags" name="item-tags" required placeholder="spicy,veg"/>
                </div>
                <div className="flex flex-row justify-between">
                    <div>Image:</div>
                    <input type="file" id="file" name="image" accept="image/*" className="file-input"/>
                </div>
                <button type="submit" className="ubuntu-bold">Create Item</button>
            </form>
        </div>
}
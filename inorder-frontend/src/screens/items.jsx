import { useEffect, useState } from "react";
import useAuthStore from "../stores/authStore";
import ItemMenu from "../components/itemMenu";
import Navbar from "../components/navbar";
import CreateItemForm from "../components/createItemForm";


export default function ItemScreen() {
    const {role,username} = useAuthStore.getState();
    const [loading, setLoading] = useState(false)

    useEffect(() => {
        if (role !== "admin") {
            window.location.href = "/notfound"
        }
    },[])



    return <div className="h-screen w-screen flex flex-col">
        <Navbar></Navbar>
        <div className="flex flex-col p-5 items-center w-full">
            <ItemMenu admin></ItemMenu>
            <div className="mt-10">
                <h1 className="text-lg font-bold mb-2">Create Item</h1>
                <CreateItemForm />
            </div>
        </div>
    </div>
}
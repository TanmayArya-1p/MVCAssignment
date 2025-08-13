import { useEffect, useState } from "react";
import useAuthStore from "@/stores/authStore";
import ItemMenu from "@/components/itemMenu";
import Navbar from "@/components/navbar";
import CreateItemForm from "@/components/createItemForm";
import Spinner from "@/components/spinner";
import VerifySignedIn from "@/utils/verify";
import { useNavigate } from "react-router-dom";
import { Toaster } from "react-hot-toast";
import { roles } from "@/utils/const";


export default function ItemScreen() {
    return <div className="h-screen w-screen flex flex-col">
        <title>Items - InOrder</title>
        <Navbar></Navbar>
        <Toaster></Toaster>
        <div className="flex flex-col p-5 items-center w-full">
            <ItemMenu admin></ItemMenu>
            <div className="mt-10">
                <div className="text-3xl font-bold mb-2">Create Item</div>
                <CreateItemForm />
            </div>
        </div>
    </div>
}
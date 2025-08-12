import { useEffect, useState } from "react";
import useAuthStore from "../stores/authStore";
import ItemMenu from "../components/itemMenu";
import Navbar from "../components/navbar";
import CreateItemForm from "../components/createItemForm";
import {fetchAllUsers, createUser, deleteUser} from "../api/users"
import UserList from "../components/userList";
import Spinner from "../components/spinner"
import CreateUserForm from "../components/createUserForm"
import { Toaster } from "react-hot-toast";

export default function UserScreen() {
    const {role,username} = useAuthStore.getState();
    const [loading, setLoading] = useState(false)
    const [users, setUsers] = useState([]);

    useEffect(() => {
        if (role !== "admin") {
            window.location.href = "/notfound"
        }
    },[])                   


    useEffect(() => {
        async function fetchUsers() {
            setLoading(true);
            const res = await fetchAllUsers();
            setUsers(res);
            setLoading(false);
        }
        fetchUsers();
    }, []);

    if(loading) {
        return <div className="h-screen w-screen flex items-center justify-center">
            <Spinner />
        </div>
    }

    return <div className="h-screen w-screen flex flex-col">
        <Navbar></Navbar>
        <Toaster />
        <div className="mt-10 p-5 flex flex-row flex-wrap gap-10 items-center justify-center">
            <UserList users={users} setUsers={setUsers}/>
            <CreateUserForm setUsers={setUsers}/>
        </div>
    </div>
}
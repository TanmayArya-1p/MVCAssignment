import { useEffect, useState } from "react";
import useAuthStore from "../stores/authStore";
import Navbar from "../components/navbar";
import {fetchAllUsers} from "../api/users"
import UserList from "../components/userList";
import Spinner from "../components/spinner"
import CreateUserForm from "../components/createUserForm"
import { Toaster } from "react-hot-toast";
import VerifySignedIn from "../utils/verify";
import { useNavigate } from "react-router-dom";
import { jwtDecode } from "jwt-decode";


export default function UserScreen() {
    const {authToken,role} = useAuthStore.getState();

    const [userID, setUserID] = useState("");

    const [loading, setLoading] = useState(false)
    const [users, setUsers] = useState([]);

    const navigate = useNavigate();

    useEffect(() => {
        if (role !== "admin") {
            navigate("/notfound");
        }
    },[])                   

    useEffect(() => {VerifySignedIn()}, [])

    useEffect(() => {
        async function fetchUsers() {
            setLoading(true);
            const res = await fetchAllUsers();
            setUsers(res);
            setLoading(false);
        }
        fetchUsers();
    }, []);


    useEffect(() => {
        setUserID(jwtDecode(authToken).userID);
    }, [authToken]);

    if(loading) {
        return <div className="h-screen w-screen flex items-center justify-center">
            <Spinner />
        </div>
    }

    return <div className="h-screen w-screen flex flex-col">
        <Navbar></Navbar>
        <title>Users - InOrder</title>
        <Toaster />
        <div className="mt-10 p-5 flex flex-row flex-wrap gap-10 items-center justify-center">
            <UserList users={users} setUsers={setUsers} authUserID={userID}/>
            <CreateUserForm setUsers={setUsers}/>
        </div>
    </div>
}
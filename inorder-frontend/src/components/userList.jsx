import {useEffect, useState} from "react"
import {deleteUser, updateUserRole} from "../api/users"
import {toast} from "react-hot-toast"
import {bumpRoleMap,bumpDownRoleMap, roles} from "../utils/const"

export default function UserList({ users,setUsers, authUserID }) {
    const [filteredUsers,setFilteredUsers] = useState(users)
    const [indexedUsers, setIndexedUsers] = useState(users)
    const [query, setQuery] = useState("")
    const [tags, setTags] = useState({
        [roles.ADMIN]: false,
        [roles.CHEF]: false,
        [roles.CUSTOMER]: false
    })


    const deleteUserHandler = async (userID) => {
        if(!window.confirm("Are you sure you want to delete this user?")) {
            return;
        }
        try {
            await deleteUser(userID)
            setUsers((prev) => prev.filter(user => user.id !== userID))
            toast.success("Successfully deleted user")

        } catch(error) {
            toast.error("Failed to delete user")
        }

    }


    const bumpDownUserHandler = async (userID) => {
        const user = users.find(user => user.id === userID);
        if(!user) return;

        if(user.role === roles.CUSTOMER) {
            toast.error("Cannot bump down customer role");
            return;
        }

        if(!window.confirm(`Are you sure you want to bump down this user's role to ${bumpDownRoleMap[user.role]}?`)) {
            return;
        }
        try {
            await updateUserRole(userID, bumpDownRoleMap[user.role]);
            setUsers((prev) => prev.map(u => u.id === userID ? {...u, role: bumpDownRoleMap[u.role]} : u));
        } catch (error) {
            toast.error("Failed to bump user role");
        }


    }

    const bumpUserHandler = async (userID) => {
        const user = users.find(user => user.id === userID);
        if(!user) return;

        if(user.role === roles.ADMIN) {
            toast.error("Cannot bump admin role");
            return;
        }

        if(!window.confirm(`Are you sure you want to bump this user's role to ${bumpRoleMap[user.role]}?`)) {
            return;
        }
        try {
            await updateUserRole(userID, bumpRoleMap[user.role]);
            setUsers((prev) => prev.map(u => u.id === userID ? {...u, role: bumpRoleMap[u.role]} : u));
        } catch (error) {
            toast.error("Failed to bump user role");
        }


    }


    useEffect(() => {
        setFilteredUsers(users);
        setIndexedUsers(users);
    }, [users])

    useEffect(() => {
        if(!(tags.admin || tags.chef || tags.customer)) {
            setFilteredUsers(users);
            return;
        }
        const filtered = users.filter(user => {
            if (tags[user.role]) {
                return true;
            }
            return false;
        });
        setFilteredUsers(filtered);
        setQuery("")
    },[tags])       



    useEffect(() => {
        const filtered = filteredUsers.filter((user) => user.username.toLowerCase().includes(query.toLowerCase()))
        setIndexedUsers(filtered)
    }, [filteredUsers, query])




    return <>
    <div className="flex flex-col gap-2 w-fit max-w-200 p-2">
        <div className="flex flex-row gap-2 items-center mt-3 bg-white border-2 rounded-sm p-2 w-full">
                <svg xmlns="http://www.w3.org/2000/svg"  fill="none" viewBox="0 0 24 24" strokeWidth={2} stroke="currentColor" className="size-6">
                    <path strokeLinecap="round" strokeLinejoin="round" d="m21 21-5.197-5.197m0 0A7.5 7.5 0 1 0 5.196 5.196a7.5 7.5 0 0 0 10.607 10.607Z" />
                </svg>
                <input type="text" id="search-input" className="w-full focus:outline-0" value={query} onChange={(e) => setQuery(e.target.value)} placeholder="Search Username"/>
        </div>
        <div id="tags-container" className="flex flex-row gap-2 items-center w-fit">
            <div className="text-lg ubuntu-bold">
                Filters:
            </div>
            {   
                Object.keys(tags).filter(k=> tags[k]).map(k=>(
                    <div key={k} className="tag tag-selected" onClick={() => setTags({...tags, [k]: false})}>{k}</div>
                ))
            }
            {   
                Object.keys(tags).filter(k=> !tags[k]).map(k=>(
                    <div key={k} className="tag" onClick={() => setTags({...tags, [k]: true})}>{k}</div>
                ))
            }


        </div>
        <div className="bg-white p-2 rounded-sm border-2 w-fit h-100 overflow-y-scroll">
            <table>
                <thead>
                    <tr className="ubuntu-bold">
                        <th className="px-4 py-2">Username</th>
                        <th className="px-4 py-2">Role</th>
                        <th className="px-4 py-2">Created At</th>
                        <th className="px-4 py-2">Delete / Change Role</th>

                    </tr>
                </thead>
                <tbody>
                    {indexedUsers.map(user => (
                        <tr key={user.id} className="border-b ubuntu-regular">
                            <td className="px-4 py-2 text-center truncate">{user.username}</td>
                            <td className="px-4 py-2 text-center ubuntu-bold truncate">{user.role}</td>
                            <td className="px-4 py-2 text-center truncate">{new Date(user.created_at).toLocaleDateString()}</td>
                            <td className="px-4 py-2 text-center flex flex-row gap-2">
                                <button className="delete-button disabled:opacity-30" disabled={user.id === authUserID} onClick={() => deleteUserHandler(user.id)}>
                                    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="size-6">
                                        <path strokeLinecap="round" strokeLinejoin="round" d="m14.74 9-.346 9m-4.788 0L9.26 9m9.968-3.21c.342.052.682.107 1.022.166m-1.022-.165L18.16 19.673a2.25 2.25 0 0 1-2.244 2.077H8.084a2.25 2.25 0 0 1-2.244-2.077L4.772 5.79m14.456 0a48.108 48.108 0 0 0-3.478-.397m-12 .562c.34-.059.68-.114 1.022-.165m0 0a48.11 48.11 0 0 1 3.478-.397m7.5 0v-.916c0-1.18-.91-2.164-2.09-2.201a51.964 51.964 0 0 0-3.32 0c-1.18.037-2.09 1.022-2.09 2.201v.916m7.5 0a48.667 48.667 0 0 0-7.5 0" />
                                    </svg>
                                </button>
                                <button className="bump-button flex flex-row gap-2 disabled:opacity-30" disabled={user.role==="admin" || user.id === authUserID} onClick={() => bumpUserHandler(user.id)}>
                                    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="size-6">
                                    <path strokeLinecap="round" strokeLinejoin="round" d="m4.5 18.75 7.5-7.5 7.5 7.5" />
                                    <path strokeLinecap="round" strokeLinejoin="round" d="m4.5 12.75 7.5-7.5 7.5 7.5" />
                                    </svg>
                                </button>
                                <button className="bump-button flex flex-row gap-2 disabled:opacity-30" disabled={user.role===roles.CUSTOMER || user.id === authUserID} onClick={() => bumpDownUserHandler(user.id)}>
                                    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="size-6">
                                    <path strokeLinecap="round" strokeLinejoin="round" d="m4.5 5.25 7.5 7.5 7.5-7.5m-15 6 7.5 7.5 7.5-7.5" />
                                    </svg>
                                </button>
                            </td>
                        </tr>
                    ))}
                </tbody>
            </table>
            {indexedUsers.length == 0 && <div className="ubuntu-thin mt-2 text-xl w-full align-middle justify-center flex">No Users Found</div>}

        </div>
    </div>
    </>
}
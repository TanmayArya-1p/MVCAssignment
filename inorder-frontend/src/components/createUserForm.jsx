import {createUser} from "../api/users"
import {toast} from 'react-hot-toast'
import { roles } from "../utils/const";

export default function CreateUserForm({setUsers}) {    

    const handleCreate = async (e) => {
        e.preventDefault();
        const formData = new FormData(e.target)
        let username = formData.get("username")
        let password = formData.get("password")
        let role = formData.get("role").toLowerCase()
        if(!(username && password && role)) {
            toast.error("Invalid parameters to create user")
            return
        }

        try {
            let resp = await createUser({username,password,role})
            setUsers((a) => [...a, resp])
            toast.success("User created successfully")
        } catch (error) {
            toast.error("Username already taken or an error occurred")    
        }
    }



    return <>
        <div id="create-user-container" className="flex flex-col gap-1 p-3 bg-white rounded-xl shadow-md border-2">
        <div className="ubuntu-bold text-2xl">Create User</div>
        <form onSubmit={handleCreate} className="space-y-4 flex flex-col gap-2 justify-center align-middle">
            <div className="flex flex-row gap-2 justify-center align-middle">
                <div>
                    <input type="text" id="username" name="username" placeholder="Username" required/>
                </div>
                <div>
                    <input type="password" id="password" name="password" placeholder="Password" required/>
                </div>
                <div>
                    <select id="role" name="role" required className="border p-2 rounded w-fit ubuntu-regular">
                        <option value="" disabled selected>Select Role</option>
                        <option value={roles.ADMIN}>Admin</option>
                        <option value={roles.CHEF}>Chef</option>
                        <option value={roles.CUSTOMER}>Customer</option>
                    </select>
                </div>
            </div>
          <button type="submit" id="create-submit" className="text-md h-fit">Create User</button>

        </form>

    </div>
    </>
}
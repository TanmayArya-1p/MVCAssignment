import AddIcon from "@/icons/addIcon";

export default function CreateOrderButton({setCreateModelOpen}) {
    return <>
        <button className="flex flex-row gap-2 text-lg items-center ubuntu-regular text-black" onClick={() => setCreateModelOpen(true)}>
            <AddIcon className="size-6" />
            <div>Create Order</div>
        </button>
    </>
}
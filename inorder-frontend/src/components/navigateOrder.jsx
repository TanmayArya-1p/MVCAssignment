import NavigateIcon from "@/icons/navigateIcon";



export default function NavigateOrder() {
    return <>
        <div id="navigate-to-order" className="ubuntu-bold mt-5 !p-2 flex flex-row gap-2 items-center order-link">
            Navigate to Order No:
            <input type="number" id="navigate-order-id" placeholder="Order No" className="p-1 border-2 w-30 mx-2 rounded-sm"/>
            <button onClick={() => navigate(`/order/${document.getElementById("navigate-order-id").value}`)}>
                <NavigateIcon className="size-5" />
            </button>
        </div>
    </>
}
export const roles = {
    ADMIN: "admin",
    CUSTOMER: "customer",
    CHEF: "chef",
}

export const orderColourMap = {
    pending: "red-500",
    preparing: "yellow-500",
    served: "teal-500",
    billed: "green-500",
    paid: "green-500",
};

export const orderItemStatusPriority = {
    pending: 3,
    preparing: 2,
    served: 1,
}

export const bumpItemStatusMap = {
    pending: "preparing",
    preparing: "served"
}

export const bumpRoleMap = {
    chef: "admin",
    customer: "chef",
}

export const modalStyle = {
    overlay: {
        backgroundColor: 'rgba(0, 0, 0, 0.25)',
    },
    content: {
    top: '50%',
    left: '50%',
    width: '60%',
    right: 'auto',
    bottom: 'auto',
    marginRight: '-50%',
    transform: 'translate(-50%, -50%)',
    border: '2px solid black',
    backgroundColor: '#fff',
    },
};


export const ONBOARDING_PATHS = ["/login", "/register", "/"];

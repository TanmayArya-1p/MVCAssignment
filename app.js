import express from "express";
import cookieParser from "cookie-parser";
import authRouter from "./routers/auth.js";
import itemsRouter from "./routers/items.js";
import usersRouter from "./routers/users.js";
import swaggerRouter from "./routers/swagger.js";
import ordersRouter from "./routers/orders.js";
import logger from "morgan";
import path from "path";
import { dirname } from "node:path";
import { fileURLToPath } from "node:url";
import * as authMiddleware from "./middleware/auth.js";
import db from "./db/index.js";
import { orderColourMap, paginate } from "./utils/misc.js";

const __dirname = dirname(fileURLToPath(import.meta.url));
const app = express();

app.set("views", path.join(__dirname, "views"));
app.set("view engine", "ejs");

app.use(logger("dev"));
app.use(express.json());
app.use(express.urlencoded({ extended: false }));
app.use(cookieParser());
app.use(express.static(path.join(__dirname, "public")));

app.get("/", function (req, res) {
  res.render("index");
});
app.get("/login", function (req, res) {
  res.render("login");
});
app.get("/register", function (req, res) {
  res.render("register");
});

app.get(
  "/home",
  authMiddleware.authenticationMiddleware(false, true),
  async function (req, res) {
    let selectedTags = req.query.tags ? req.query.tags.split(",") : [];
    let selectHM = {};
    for (const i of selectedTags) {
      selectHM[i] = true;
    }

    let orders = null;
    if (res.locals.user.role === "customer") {
      orders = await db.Order.getAllOrdersByUser(res.locals.user);
    } else {
      orders = await db.Order.getAllOrders();
    }

    let items = null;
    let page = null;
    switch (res.locals.user.role) {
      case "customer":
        if (selectedTags.length > 0) {
          items = await db.Item.getItemsofTag(selectedTags, -1, 0);
        } else {
          items = await db.Item.getAllItems(-1, 0);
        }
        page = paginate(items, req);
        items = page.filtered;
        res.render(`customer-home`, {
          user: res.locals.user,
          orders: orders,
          items: items,
          orderColourMap: orderColourMap,
          tags: await db.Tags.getAllTags(),
          selectedTags: selectHM,
          page: page,
        });
        break;
      case "chef":
        page = paginate(orders, req);
        orders = page.filtered;

        items = await db.Item.getAllItems(-1, 0);
        let itemHM = {};
        for (const item of items) {
          itemHM[item.id] = item;
        }

        let orderedItems = await db.OrderItems.getAllOrderedItems();
        orderedItems = orderedItems.filter(
          (a) => a.status === "pending" || a.status === "preparing",
        );
        orderedItems.sort((a, b) => {
          if (a.status === b.status) return 0;
          if (a.status === "preparing") return -1;
          if (b.status === "preparing") return 1;
          return 0;
        });

        res.render(`chef-home`, {
          user: res.locals.user,
          orders: orders,
          orderedItems: orderedItems,
          orderColourMap: orderColourMap,
          page: page,
          itemHM: itemHM,
        });
        break;

      case "admin":
        page = paginate(orders, req);
        orders = page.filtered;

        items = await db.Item.getAllItems(-1, 0);

        res.render(`admin-home`, {
          user: res.locals.user,
          orders: orders,
          orderColourMap: orderColourMap,
          page: page,
          items: items,
        });
        break;
    }
  },
);

app.get(
  "/order/create",
  authMiddleware.authenticationMiddleware(false, true),
  async function (req, res) {
    res.render("create-order", {
      user: res.locals.user,
      items: await db.Item.getAllItems(),
      tags: await db.Tags.getAllTags(),
    });
  },
);

app.get(
  "/order/:orderid",
  authMiddleware.authenticationMiddleware(false, true),
  async function (req, res) {
    let order = await db.Order.getOrderById(Number(req.params.orderid));

    if (!order) {
      return res.status(404).send("Not Found");
    }

    if (
      res.locals.user.role === "customer" &&
      order.issued_by !== res.locals.user.id
    ) {
      return res.status(403).send("Forbidden");
    }

    res.render("order-view", {
      user: res.locals.user,
      items: await db.Item.getAllItems(),
      order: order,
      orderColourMap: orderColourMap,
      tags: await db.Tags.getAllTags(),
    });
  },
);

app.use("/api/auth", authRouter);
app.use("/api/items", itemsRouter);
app.use("/api/users", usersRouter);
app.use("/api/orders", ordersRouter);
app.use("/api/swagger", swaggerRouter);

app.use(function (req, res, next) {
  res.status(404).send("404 Not Found");
});

export default app;

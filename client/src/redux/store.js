// import { configureStore } from "@reduxjs/toolkit";
import IoTReducer from "./IoTSlice";
import authReducer from "./authSlice";
// const store = configureStore({
//   reducer: {
//     IoT: IoTReducer
//   }
// });

// export default store;

import { configureStore, combineReducers } from "@reduxjs/toolkit";
import {
    persistStore,
    persistReducer,
    FLUSH,
    REHYDRATE,
    PAUSE,
    PERSIST,
    PURGE,
    REGISTER,
} from "redux-persist";
import storage from "redux-persist/lib/storage";
const persistConfig = {
    key: "root",
    version: 1,
    storage,
};
const rootReducer = combineReducers({ IoT: IoTReducer, auth_: authReducer})
const persistedReducer = persistReducer(persistConfig, rootReducer);
export const store = configureStore({
    reducer: persistedReducer,
    middleware: (getDefaultMiddleware) =>
        getDefaultMiddleware({
            serializableCheck: {
                ignoredActions: [FLUSH, REHYDRATE, PAUSE, PERSIST, PURGE, REGISTER],
            },
        }),
})
export let persistor = persistStore(store);

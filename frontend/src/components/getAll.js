import React from "react";
import axios  from "axios";
import { API_URL } from "../config";

const GetUrl = API_URL + "/inventoryList"


export default function GetItems(){
    const [itemsList, setItemsList] = React.useState(null);
    React.useEffect(
        () => {
            axios.get(GetUrl)
                .then((response) => {
                    console.table(response.data)
                    setItemsList(response.data)
                });
        }, 
        []);
    if (!itemsList) return null;
    return(
        <>
            <div>
                <h1>Inventory Table</h1>
            </div>
            <table>
                <tr>
                    <th>ID</th>
                    <th>Name</th>
                    <th>Price</th>
                    <th>Sales</th>
                    <th>Stock</th>
                </tr>
                {itemsList.map((eachItem, index) => {
                    return (
                        <tr key={index}>
                            <td>{eachItem.id}</td>
                            <td>{eachItem.name}</td>
                            <td>{eachItem.price}</td>
                            <td>{eachItem.sales}</td>
                            <td>{eachItem.stock}</td>
                        </tr>
                    )
                } )}
            </table>
        </>
    )
}
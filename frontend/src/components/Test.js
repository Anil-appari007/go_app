import axios from 'axios'
import React, {useEffect, useState} from 'react'
import { API_URL } from '../config'

const URL = API_URL + "/inventoryList"

const ADD_URL = API_URL + "/addItem"
const PUT_URL = API_URL + "/updateItem"
export default function Table() {
    if (!API_URL) {
        console.error('Error: env var REACT_APP_API_URL found to be empty');
    } else {
        console.log('Found env REACT_APP_API_URL', API_URL)
    }
    const [items, setItems] = useState([])
    const [id, setId] = useState('')
    const [name, setName] = useState('')
    const [price, setPrice] = useState('')
    const [sales, setSales] = useState('')
    const [stock, setStock] = useState('')

    const [uprice, setUprice] = useState('')
    const [usales, setUsales] = useState('')
    const [ustock, setUstock] = useState('')

    // const [editName, setEditName] = useState('')
    const [editId, setEditId] = useState('')
    const [editName, setEditName] = useState('')

    useEffect(() => {
        axios.get(URL)
        // .then(response => console.log(response))
        .then(response => setItems(response.data))
        .catch(error => console.log(error));
    }, [])

    const handleSubmit = (event) => {
        event.preventDefault();
        // console.log({"id":id,"name":name,"price":price,"sales":sales,"stock":stock})


        const iId= parseInt(id, 10);
        const iPrice = parseInt(price, 10);
        const iSales = parseInt(sales, 10);
        const iStock = parseInt(stock, 10);
        console.log({"id":iId,"name":name,"price":iPrice,"sales":iSales,"stock":iStock})
        const postData = {
            "id": iId,
            "name": name,
            "price": iPrice,
            "sales": iSales,
            "stock": iStock,
        };
        axios.post(ADD_URL, postData, {
            headers: {'Content-Type': 'application/json',},
        })
            .then(response => {
                // console.log(response)
                window.location.reload();
            } )
            .catch(err => console.log(err));
    }

    const handleEdit = (id, name) => {
        axios.get(URL + '/' + name)
        .then(response => {
            console.log(response)
            setUprice(response.data.price)
            setUsales(response.data.sales)
            setUstock(response.data.stock)
        })
        .catch(err => console.log(err));
        setEditId(id)
        setEditName(name)
    }

    const handleUpdate = () => {
        const iuPrice = parseInt(uprice, 10);
        const iuSales = parseInt(usales, 10);
        const iuStock = parseInt(ustock, 10);
        // console.log({"id":editId,"name":name,"price":iuPrice,"sales":iuSales,"stock":iuStock})
        const postData = {
            "id": editId,
            "name": editName,
            "price": iuPrice,
            "sales": iuSales,
            "stock": iuStock,
        };
        console.log(PUT_URL,postData)
        axios.put(PUT_URL, postData, {
            headers: {'Content-Type': 'application/json',},
        })
            .then(response => {
                console.log(response)
                window.location.reload();
            } )
            .catch(err => console.log(err));
    }

    const handleDelete = (item) => {
        const postData = {
            "id": item.id,
            "name": item.name,
            "price": item.price,
            "sales": item.sales,
            "stock": item.stock,
        };
        console.log(postData)
        axios.delete(API_URL + '/deleteItem', {
            data: postData,
            headers: { 'Content-Type': 'application/json' },
        })
        // .then(response => console.log(response))
        .then(response => {
            window.location.reload();
            // console.log(response.data)
        })
        .catch(err => console.log(err));
    }
    return(
        <div className='container'>
            <div>
                <form onSubmit={handleSubmit}>
                    <input type='id' placeholder='Enter Product ID' onChange={e => setId(e.target.value)}></input>
                    <input type='Name' placeholder='Enter Product Name' onChange={e => setName(e.target.value)}></input>
                    <input type='Price' placeholder='Enter Product Price' onChange={e=> setPrice(e.target.value)}></input>
                    <input type='Sales' placeholder='Enter Product Sales' onChange={e=> setSales(e.target.value)}></input>
                    <input type='Stock' placeholder='Enter Product Stock' onChange={e => setStock(e.target.value)}></input>
                    <button>Add</button>
                </form>
            </div>
            <table>
                <thead>
                    <tr>
                        <th>ID</th>
                        <th>Name</th>
                        <th>Price</th>
                        <th>Sales</th>
                        <th>Stock</th>
                        <th>Action</th>
                    </tr>
                </thead>
                <tbody>
                    {
                        items.map((item, index) => (
                            item.id === editId ?
                            <tr key={index}>
                                <td> {item.id} </td>
                                <td> {item.name} </td>
                                <td> <input type='text' value={item.uprice} onChange={e => setUprice(e.target.value)}/> </td>
                                <td> <input type='text' value={item.usales} onChange={e => setUsales(e.target.value)}/> </td>
                                <td> <input type='text' value={item.ustock} onChange={e => setUstock(e.target.value)}/> </td>
                                <td> <button onClick={handleUpdate}>Update</button> </td>
                            </tr>
                            :
                            <tr key={index}>
                                <td> {item.id} </td>
                                <td> {item.name} </td>
                                <td> {item.price} </td>
                                <td> {item.sales} </td>
                                <td> {item.stock} </td>
                                <td>
                                    <button onClick={() => handleEdit(item.id, item.name)}>edit</button>
                                    <button onClick={() => handleDelete(item)}>delete</button>
                                </td>
                            </tr>
                         ) )
                    }
                </tbody>
            </table>
        </div>
    )
}
import React, {useState} from 'react'
import {Button, Form} from 'semantic-ui-react'
import axios from 'axios';
export default function Create() {
    const [name, setName] = useState('');
    const [id, setId] = useState('');
    const [price, setPrice] = useState('');
    const [sales, setSales] = useState('');
    const [stock, setStock] = useState('');
    const postData = () => {

        console.log(name);
        console.log(id);
        console.log(price);
        console.log(sales);
        console.log(stock);
        axios.post(
            'http://localhost:8888/addItem', 
            {
                "id":id,
                "name": name,
                "price": price,
                "sales": sales,
                "stock": stock
            }
        )
    }
    return(
        <Form className='create-form'>
            <Form.Field >
                <label>Item Name</label>
                <input placeholder='Item Name' onChange={(e) => setName(e.target.value)}></input>
            </Form.Field>
            <Form.Field>
                <label>Item ID</label>
                <input placeholder='Item ID' onChange={(e) => setId(e.target.value)}></input>
            </Form.Field>
            <Form.Field>
                <label>Price</label>
                <input placeholder='Price' onChange={(e) => setPrice(e.target.value)}></input>
            </Form.Field>
            <Form.Field>
                <label>Item Sales</label>
                <input placeholder='Item Sales' onChange={(e) => setSales(e.target.value)}></input>
            </Form.Field>
            <Form.Field>
                <label>Item Stock</label>
                <input placeholder='Item Stock' onChange={(e) => setStock(e.target.value)}></input>
            </Form.Field>
            <Button onClick={postData} type='submit'>submit</Button>
        </Form>
    )
}

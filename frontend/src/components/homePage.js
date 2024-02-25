import React from 'react'
import axios from "axios";
import { API_URL } from "../config";

const GetUrl = API_URL + "/hello"
export default function HomePage(){
    const [post, setPost] = React.useState(null);

    React.useEffect(() => {
        axios.get(GetUrl).then(
            (response) => {setPost(response.data)}
        )
    }, []
    );
    if (!post) return null;
    return(
        <div>
            <h1>Message from {post.message}</h1>
        </div>
    )
}
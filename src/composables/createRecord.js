import { ref } from 'vue'
import axios from 'axios'
import { useRouter } from 'vue-router'

const createRecord = () => {
    const router = useRouter()
    const file = ref(null)
    const fileUrl = ref('')
    const at = localStorage.getItem('access_token')
    const access_token = {
        headers: { Authorization: `Bearer ${at}` }
    };

    const writeToDb = async (query, ...data) => {
        axios
            .post("/api/createRecord", {
                query: query,
                data: data,
            })
            .then(response => {
                router.go()
            })
    }

    const handleChange = (e) => {
        // allowed file types
        const types = ['image/png', 'image/jpeg']

        // selected file
        const selected = e.target.files[0]

        if (selected && types.includes(selected.type)) {
            file.value = selected
        } else {
            file.value = null
        }
    }

    const uploadImage = async () => {
        let formData = new FormData();
        
        formData.append("file", file.value);
        axios.post('/api/upload', formData, access_token)
        .then(res => {
            console.log(res)
        })
        .catch(err => {
            console.log(err)
        })

        // fetch('http://localhost:3090/api/upload', {
        //     method: 'post',
        //     body: formData,
        // }).then(res => res.json())
        // .then(data => {
        //     console.log(data.url)
        //     fileUrl.value = data.url
        // })
        // .catch(err => {
        //     console.log(err)
        // })
    }

    return { writeToDb, handleChange, uploadImage, fileUrl }
}

export default createRecord

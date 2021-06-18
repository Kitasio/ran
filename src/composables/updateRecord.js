import { ref } from 'vue'
import axios from 'axios'
import { useRouter } from 'vue-router'

const updateRecord = () => {
    const router = useRouter()
    const file = ref(null)
    const fileName = ref('')
    const filePath = ref('')
    const fileUrl = ref('')

    const updateToDb = async (query, ...data) => {
        console.log(query, data)
        axios
            .post("http://localhost:3090/api/updateRecord", {
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
            let formData = new FormData();
                
            formData.append("file", file.value);
            axios.post('http://localhost:3090/api/upload', formData).then(r => {
                fileUrl.value = '/' + r.data
                console.log(fileUrl.value)
            })
        } else {
            file.value = null
        }
    }


    return { updateToDb, handleChange, fileUrl }
}

export default updateRecord
import axios from 'axios'

const ossGetKey = (form) => axios.post("/api/v1/oss/key", form).then(response => response.data);

const replyGet = (id) => axios.get(`/api/v1/reply/${id}`).then(response => response.data);

const replyCreate = (form) => axios.post("/api/v1/reply", form).then(response => response.data);

const userLogin = (form) => axios.post('api/v1/user/login', form).then(response => response.data);

export {
  ossGetKey,
  replyGet,
  replyCreate,
  userLogin

}

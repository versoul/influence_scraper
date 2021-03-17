import axios from 'axios'

// eslint-disable-next-line
export const CheckRate = async function (nickname: string) {
  const resp = await axios.get('/api/check-rate/' + nickname)
  if (resp.data.code !== 200) {
    throw new Error(resp.data.error)
  }

  return resp.data
}

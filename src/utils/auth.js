import Cookies from 'js-cookie'

const TokenKey = 'ginx_github.com/karry-11110/LM_system'

export function getToken() {
  return Cookies.get(TokenKey)
}

export function setToken(token) {
  return Cookies.set(TokenKey, token)
}

export function removeToken() {
  return Cookies.remove(TokenKey)
}

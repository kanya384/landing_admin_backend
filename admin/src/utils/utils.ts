import Cookies from 'js-cookie';

export const GetTokenFromCookies = (): string => {
  let token = ""
  let tmp = Cookies.get("token")
  if (tmp) {
    token = tmp
  }
  return token
}
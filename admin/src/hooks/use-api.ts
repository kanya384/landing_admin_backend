import { useMemo } from "react";
import { DefaultApi, Configuration, PostersApi } from "../api";

export const useApi = () => {
  const url = "http://localhost:3000/"
  const authService = new DefaultApi(new Configuration(), url);
  const postersService = new PostersApi(new Configuration(), url)
  return useMemo(()=>{
    return {
      authService,
      postersService
    }
  }, [authService, postersService])
}
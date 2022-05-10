import { Doc } from "../../api";
import { DocsActionTypes } from "../action-types";

export interface DocsRequestSend{
  type: DocsActionTypes.DOCS_REQUEST_SEND
}

export interface DocsError {
  type: DocsActionTypes.DOCS_ERROR,
  payload: string,
}

export interface DocsNew {
  type: DocsActionTypes.DOCS_NEW,
  payload: Doc
}

export interface DocsSuccess {
  type: DocsActionTypes.DOCS_SUCCESS,
  payload: Doc[],
}

export interface DocsUpdate {
  type: DocsActionTypes.DOCS_UPDATE,
  payload: Doc,
}

export interface DocsDelete {
  type: DocsActionTypes.DOCS_DELETE,
  payload: string,
}

export type DocAction =
  | DocsRequestSend
  | DocsError
  | DocsSuccess
  | DocsNew
  | DocsUpdate
  | DocsDelete
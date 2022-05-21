import { FC } from "react";
import { Lead } from "../../api";

interface TableItemProps {
  lead: Lead,
  index: number,
  deleteAction: any
}

export const TableItem: FC<TableItemProps> = (props) => {
  return(<tr>
          <th scope="row">{props.index}</th>
          <td>{new Date(Date.parse(props.lead.createdAt!)).toISOString().slice(0,10)}</td>
          <td>{props.lead.phone}</td>
          <td>{props.lead.name}</td>
          <td>{props.lead.email}</td>
          <td>{props.lead.text}</td>
          <td className="col-sm" style={{width:"235px"}}>
            <button className="btn btn-soft-danger btn-sm" type="button" onClick={()=>{props.deleteAction()}}>Удалить</button>
          </td>
        </tr>
  )
}
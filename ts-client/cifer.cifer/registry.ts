import { GeneratedType } from "@cosmjs/proto-signing";
import { QueryParamsRequest } from "./types/cifer/cifer/query";
import { QueryParamsResponse } from "./types/cifer/cifer/query";
import { MsgUpdateParams } from "./types/cifer/cifer/tx";
import { MsgUpdateParamsResponse } from "./types/cifer/cifer/tx";
import { GenesisState } from "./types/cifer/cifer/genesis";
import { Params } from "./types/cifer/cifer/params";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/cifer.cifer.QueryParamsRequest", QueryParamsRequest],
    ["/cifer.cifer.QueryParamsResponse", QueryParamsResponse],
    ["/cifer.cifer.MsgUpdateParams", MsgUpdateParams],
    ["/cifer.cifer.MsgUpdateParamsResponse", MsgUpdateParamsResponse],
    ["/cifer.cifer.GenesisState", GenesisState],
    ["/cifer.cifer.Params", Params],
    
];

export { msgTypes }
export interface CodebasePayload {
    codebase_id: string;
    codebase_name: string;
    github_url: string;
    username: string;
    token: string;
    branch: string;
    folder_path: string;
}



export interface ExtractCodePayload {
    codebase_name: string;
    github_url: string;
    username: string;
    token: string;
    branch: string;
    folder_path: string;
}


export interface IUserChatSessionsData {
    session_id: string;
    codebase_ids: string[];
}

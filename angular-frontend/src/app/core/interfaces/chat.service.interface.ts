export interface IChat {
    session_id: string
    chat_id: string;
    user_id: string;
    ai_answer: string;
    user_question: string;
    references: ReferencesWithSimilarity[]
}



export interface ReferencesWithSimilarity {
    vector_id: string;
    codebase_id: string;
    codebase_name: string;
    hashId: string;
    filePath: string;
    code: string;
    embedding: number[];
    similarity_to_query: number;
    similarity_to_prev_query: number;
}


type StoryItem = {
    id_story: number, 
    link_story: string,
    dt_criacao: string,
    visualizou: boolean,
}
type StoryResponse = {
    id_usuario: number,
    nome: string,
    link_perfil: string,
    arroba: string,
    stories: string,
}

type Story = Omit<StoryResponse, 'stories'> & {
    stories: StoryItem[],
}
export type {
    StoryResponse,
    StoryItem,
    Story
}
const timeAgo = (date: string) => {
    if(new Date(date) > new Date(Date.now())){
        return 'No futuro'
    }

    const now = new Date()
    const postDate = new Date(date)
    const diff = now.getTime() - postDate.getTime()
    const diffYears = Math.floor(diff / (1000 * 60 * 60 * 24 * 7 * 4 * 12))
    const diffWeeks = Math.floor(diff / (1000 * 60 * 60 * 24 * 7))
    const diffDays = Math.floor(diff / (1000 * 60 * 60 * 24))
    const diffHours = Math.floor(diff / (1000 * 60 * 60))
    const diffMinutes = Math.floor(diff / (1000 * 60))
    const diffSeconds = Math.floor(diff / (1000))
    if(diffYears > 0){
        return diffYears === 1 ? '1ano atrás' : diffYears + 'anos atrás'
    }else if(diffWeeks > 0){
        return diffWeeks === 1 ? '1s atrás' : diffWeeks + 's atrás'
    }else if(diffDays > 0){
        return diffDays === 1 ? '1d atrás' : diffDays + 'd atrás'
    }else if(diffHours > 0){
        return diffHours === 1 ? '1h atrás' : diffHours + 'h atrás'
    }else if(diffMinutes > 0){
        return diffMinutes === 1 ? '1m atrás' : diffMinutes + 'm atrás'
    }else{
        return Math.abs(diffSeconds) + 's atrás'
    }
}

export {
    timeAgo
}
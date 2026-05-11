import { useEffect, useRef, useState } from 'react'

export interface GalleryCard {
  contributor_id: string
  html_content: string
  approved_at: string
  percentage: number
}

export function useGalleryWebSocket(url: string) {
  const [cards, setCards] = useState<GalleryCard[]>([])
  const [isConnected, setIsConnected] = useState(false)
  const ws = useRef<WebSocket | null>(null)

  useEffect(() => {
    let reconnectTimeout: NodeJS.Timeout

    const connect = () => {
      ws.current = new WebSocket(url)

      ws.current.onopen = () => {
        setIsConnected(true)
        console.log('Connected to Gallery WebSocket')
      }

      ws.current.onmessage = (event) => {
        try {
          const lines = event.data.split('\n').filter((line: string) => line.trim() !== '')

          setCards((prev) => {
            const cardMap = new Map<string, GalleryCard>()
            
            // Initialize map with current cards
            prev.forEach(c => cardMap.set(c.contributor_id, c))
            
            // Process incoming lines
            for (const line of lines) {
              try {
                const card: GalleryCard = JSON.parse(line)
                const existing = cardMap.get(card.contributor_id)
                
                // Only keep the card if it's new or more recent than what we have
                if (!existing || new Date(card.approved_at) > new Date(existing.approved_at)) {
                  cardMap.set(card.contributor_id, card)
                }
              } catch (err) {
                console.error('Error parsing individual gallery line:', err)
              }
            }
            
            // Return sorted array (most recent first)
            return Array.from(cardMap.values()).sort((a, b) => 
              new Date(b.approved_at).getTime() - new Date(a.approved_at).getTime()
            )
          })
        } catch (err) {
          console.error('Error parsing gallery message:', err)
        }
      }

      ws.current.onclose = () => {
        setIsConnected(false)
        console.log('Disconnected from Gallery WebSocket. Reconnecting in 5s...')
        reconnectTimeout = setTimeout(connect, 5000)
      }

      ws.current.onerror = (err) => {
        console.error('WebSocket error:', err)
        ws.current?.close()
      }
    }

    connect()

    return () => {
      clearTimeout(reconnectTimeout)
      if (ws.current) {
        // Prevent onclose from triggering reconnect during unmount
        ws.current.onclose = null
        ws.current.close()
      }
    }
  }, [url])

  return { cards, isConnected }
}

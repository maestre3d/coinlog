import './globals.css'

export const metadata = {
    title: 'Coinlog',
    description: 'Coinlog is an assistant to keep track of your personal finance records.',
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en">
      <body>{children}</body>
    </html>
  )
}

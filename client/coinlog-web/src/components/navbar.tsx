import Link from 'next/link'
import { Poppins } from 'next/font/google'

const poppins = Poppins({ weight: ['700'], subsets: ['latin'] })

export default function Navbar() {
    return (
        <div className='flex w-full flex-row justify-between'>
            <h1 className={`text-3xl font-bold ${poppins.className}`}>Coinlog</h1>
            <Link href='/auth' className='text-bold text-3xl'>Sign In</Link>
        </div>
    )
}
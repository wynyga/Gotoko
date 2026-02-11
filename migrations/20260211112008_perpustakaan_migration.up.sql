-- 1. Tabel Users
CREATE TABLE IF NOT EXISTS public.users (
    id character varying(36) DEFAULT gen_random_uuid() NOT NULL,
    email character varying(255) NOT NULL,
    password character varying(255) NOT NULL,
    CONSTRAINT users_pk PRIMARY KEY (id)
);

-- 2. Tabel Books
CREATE TABLE IF NOT EXISTS public.books (
    id character varying(36) DEFAULT gen_random_uuid() NOT NULL,
    title character varying(255) NOT NULL,
    description text,
    isbn character varying(100) NOT NULL,
    created_at timestamp(6) without time zone,
    updated_at timestamp(6) without time zone,
    deleted_at timestamp(6) without time zone,
    CONSTRAINT books_pk PRIMARY KEY (id)
);

-- 3. Tabel Customers
CREATE TABLE IF NOT EXISTS public.customers (
    id character varying(36) DEFAULT gen_random_uuid() NOT NULL,
    code character varying(255) NOT NULL,
    name character varying(255) NOT NULL,
    created_at timestamp(6) without time zone,
    updated_at timestamp(6) without time zone,
    deleted_at timestamp(6) without time zone,
    CONSTRAINT customers_pk PRIMARY KEY (id)
);

-- 4. Tabel Book Stocks
CREATE TABLE IF NOT EXISTS public.book_stocks (
    book_id character varying(36) NOT NULL,
    code character varying(50) NOT NULL,
    status character varying(50) NOT NULL,
    borrower_id character varying(36),
    borrowed_at timestamp(6) without time zone,
    CONSTRAINT book_stocks_pk PRIMARY KEY (code)
);

-- Optional: Tambahkan Foreign Key jika diperlukan untuk relasi (Good practice di Go Microservices)
-- ALTER TABLE public.book_stocks ADD CONSTRAINT fk_book_stocks_book FOREIGN KEY (book_id) REFERENCES public.books(id);
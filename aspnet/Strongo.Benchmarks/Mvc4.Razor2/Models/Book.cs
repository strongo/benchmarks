using System;
using System.Collections.Generic;
using System.Linq;
using System.Web;

namespace Mvc4.Razor2.Models
{
    public class Book
    {
        public int Id { get; set; }
        public string Title { get; set; }
        public DateTime FirstTimePublished { get; set; }
        public int Pages { get; set; }
        public decimal Price { get; set; }
    }
}
CREATE MIGRATION m1a7ijfjf35xwfrby66mycucvaggldsr6szhyngy2jfu5qukw53pvq
    ONTO initial
{
  CREATE TYPE default::User {
      CREATE PROPERTY password -> std::str;
      CREATE PROPERTY dob -> std::str;
      CREATE PROPERTY verified -> std::bool;
      CREATE PROPERTY firstname -> std::str;
      CREATE PROPERTY lastname -> std::str;
      CREATE PROPERTY phone -> std::str;
      CREATE PROPERTY cv -> std::str;
      CREATE PROPERTY accesstoken -> std::str;
      CREATE PROPERTY refreshtoken -> std::str;
      CREATE REQUIRED PROPERTY email -> std::str {
          CREATE CONSTRAINT std::exclusive;
      };
  };

};

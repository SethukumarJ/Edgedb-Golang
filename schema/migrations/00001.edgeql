CREATE MIGRATION m1nztvmah5qfwdshtzf5ojqv34ffoamqro6bmtbnjcu4jxzyxcs7sa
    ONTO initial
{
  CREATE TYPE default::User {
      CREATE REQUIRED PROPERTY password -> std::str;
      CREATE REQUIRED PROPERTY dob -> std::str;
      CREATE REQUIRED PROPERTY verified -> std::bool;
      CREATE REQUIRED PROPERTY firstname -> std::str;
      CREATE REQUIRED PROPERTY lastname -> std::str;
      CREATE REQUIRED PROPERTY phone -> std::str;
      CREATE REQUIRED PROPERTY cv -> std::str;
      CREATE PROPERTY accesstoken -> std::str;
      CREATE PROPERTY refreshtoken -> std::str;
      CREATE REQUIRED PROPERTY email -> std::str {
          CREATE CONSTRAINT std::exclusive;
      };
  };

};
